package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"testing"
)

func testFind(t *testing.T, filenames []string, search *regexp.Regexp, want []string) {
	t.Helper()
	tempDir, err := ioutil.TempDir("", "")
	defer os.RemoveAll(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, filename := range filenames {
		file := filepath.Join(tempDir, filename)
		if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
			t.Fatal(err)
		}
		err := ioutil.WriteFile(file, []byte("foobar"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	finder := NewFinder(search)
	got, err := finder.Find(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	var absWant []string
	for _, w := range want {
		absWant = append(absWant, filepath.Join(tempDir, w))
	}

	sort.Strings(got)
	sort.Strings(absWant)

	if !reflect.DeepEqual(got, absWant) {
		t.Errorf("got: %v\n want: %v\n", got, absWant)
	}
}

func Test_FindSingleMatch(t *testing.T) {
	filenames := []string{"src/foo/foo.go", "src/bar/bar.go", "src/duck/feathered/feathered.go"}
	search, _ := regexp.Compile(filenames[0])
	want := []string{filenames[0]}
	testFind(t, filenames, search, want)
}

func Test_FindMultipleMatches(t *testing.T) {
	filenames := []string{"src/foo/foo.go", "src/bar/baz/foo.go", "src/duck/feathered/feathered.go"}
	search, _ := regexp.Compile("foo.go")
	want := []string{filenames[0], filenames[1]}
	testFind(t, filenames, search, want)
}

func Test_FindNoMatches(t *testing.T) {
	filenames := []string{"src/foo/foo.go", "src/bar/baz/baz.go", "src/duck/feathered/feathered.go"}
	search, _ := regexp.Compile("xyz.go")
	var want []string
	testFind(t, filenames, search, want)
}
