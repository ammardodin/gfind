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

type void struct{}

type Finder struct {
	startDir string
	regex    regexp.Regexp
	work     chan string // from main thread to workers
	workFeed chan string // from workers to main thread
	errors   chan error  // from workers to main thread
	results  []string    // from workers to main thread
	done     chan void   // to terminate workers gracefully
	mutex    sync.Mutex
}

func NewFinder(startDir string, regex regexp.Regexp) *Finder {
	numWorkers := runtime.NumCPU()
	f := &Finder{
		startDir: startDir,
		regex:    regex,
		work:     make(chan string, numWorkers),
		workFeed: make(chan string, numWorkers),
		errors:   make(chan error, numWorkers),
		done:     make(chan void),
	}
	return f
}

func (f *Finder) find(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			f.workFeed <- filePath
		} else if f.regex.MatchString(filePath) {
			f.mutex.Lock()
			f.results = append(f.results, filePath)
			f.mutex.Unlock()
		}
	}

	return nil
}

func (f *Finder) waitAndWork(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-f.done:
			return
		case path := <-f.work:
			f.errors <- f.find(path)
		}
	}
}

func (f *Finder) Find() ([]string, error) {
	wg := &sync.WaitGroup{}

	defer close(f.errors)
	defer close(f.workFeed)
	defer close(f.work)
	defer wg.Wait()
	defer close(f.done)

	numWorkers := runtime.NumCPU()
	fmt.Printf("Using %d workers !\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go f.waitAndWork(wg)
	}

	queue := StringQueue{}
	queue.Push(f.startDir)
	inflight := 0

	for {
		work := f.work
		var path string
		var err error

		if queue.Empty() {
			// Disable first case statement when queue is empty
			work = nil
		} else {
			path, err = queue.Front()
			if err != nil {
				return nil, err
			}
		}

		select {
		case work <- path:
			queue.Pop()
			inflight++
		case path := <-f.workFeed:
			queue.Push(path)
		case err = <-f.errors:
			inflight--
			if err != nil {
				fmt.Fprintf(os.Stderr, Error(err))
			}
			if inflight == 0 && queue.Empty() {
				select {
				case path := <-f.workFeed:
					queue.Push(path)
				default:
					return f.results, nil
				}
			}
		}
	}
}
