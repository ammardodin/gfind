/**
Lifted from https://gist.github.com/ik5/d8ecde700972d4378d87#gistcomment-3074524
 */
package main

import "fmt"

var (
	Info = Teal
	Error = Red
)

var (
	Red     = Color("\033[1;31m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

