package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	rawStart, rawExpr := os.Args[1], os.Args[2]

	absStart, err := filepath.Abs(rawStart)
	if err != nil {
		log.Fatal(err)
	}
	cleanStart := filepath.Clean(absStart)

	expr, err := regexp.Compile(rawExpr)
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
