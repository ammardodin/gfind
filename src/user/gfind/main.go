// Package gfind implements a concurrent file finder.

package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
)

func main() {
	rawStart := flag.String("start", ".", "Absolute or relative starting path for the search")
	rawExpr := flag.String("expr", "*", "Search expression to match files against")
	flag.Parse()

	absStart, err := filepath.Abs(*rawStart)
	if err != nil {
		log.Fatal(err)
	}
	cleanStart := filepath.Clean(absStart)

	expr, err := regexp.Compile(*rawExpr)
	if err != nil {
		log.Fatal(err)
	}

	finder := NewFinder(cleanStart, *expr)
	fmt.Printf("start: %s, expr: %s\n", absStart, expr)
	matches, err := finder.Find()
	for _, match := range matches {
		fmt.Println(match)
	}
}
