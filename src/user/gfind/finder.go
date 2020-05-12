package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
)

type Finder struct {
	regex      *regexp.Regexp
	work       chan string // from main thread to workers
	workFeed   chan string // from workers to main thread
	errors     chan error  // from workers to main thread
	results    []string    // from workers to main thread
	done       chan byte   // to terminate workers gracefully
	dispatched int         // counter for inflight work
	mutex      sync.Mutex
}

func NewFinder(regex *regexp.Regexp) *Finder {
	numWorkers := runtime.NumCPU()
	return &Finder{
		regex:    regex,
		work:     make(chan string, numWorkers),
		workFeed: make(chan string, numWorkers),
		errors:   make(chan error, numWorkers),
		done:     make(chan byte),
	}
}

func (finder *Finder) find(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			finder.workFeed <- filePath
		} else if finder.regex.MatchString(filePath) {
			finder.mutex.Lock()
			finder.results = append(finder.results, filePath)
			finder.mutex.Unlock()
		}
	}

	return nil
}

func (finder *Finder) worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-finder.done:
			return
		case dir := <-finder.work:
			finder.errors <- finder.find(dir)
		}
	}
}

func (finder *Finder) Find(startDir string) ([]string, error) {
	wg := &sync.WaitGroup{}

	defer close(finder.errors)
	defer close(finder.workFeed)
	defer close(finder.work)
	defer wg.Wait()
	defer close(finder.done)

	numWorkers := runtime.NumCPU()
	fmt.Printf("Using %d workers !\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go finder.worker(wg)
	}

	queue := StringQueue{}
	queue.Push(startDir)

	for {
		work := finder.work
		var path string
		var err error

		if queue.Empty() {
			// Disable second case statement when queue is empty
			work = nil
		} else {
			path, err = queue.Front()
			if err != nil {
				return nil, err
			}
		}

		select {
		case path := <-finder.workFeed:
			queue.Push(path)
		case work <- path:
			_, err = queue.Pop()
			if err != nil {
				return nil, err
			}
			finder.dispatched++
		case err = <-finder.errors:
			finder.dispatched--
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprint(err))
			}
			if finder.dispatched == 0 && queue.Empty() {
				select {
				case path := <-finder.workFeed:
					queue.Push(path)
				default:
					finder.done <- byte(1)
					return finder.results, nil
				}
			}
		}
	}
}
