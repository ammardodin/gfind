// largely lifted from https://eli.thegreenplace.net/2020/testing-flag-parsing-in-go-programs/
package main

import (
	"flag"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

type parseResult struct {
	conf *config
	err  error
}

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args []string
		want parseResult
	}{
		{[]string{"-start", "/usr/local"},
			parseResult{nil, MissingPatternErr},
		},
		{[]string{"-pattern", "foo.baz"},
			parseResult{nil, MissingStartErr},
		},
		{[]string{"-start", "/", "-pattern", "foo.go"},
			parseResult{&config{start: "/", pattern: regexp.MustCompile("foo.go")}, nil},
		},
		{
			[]string{"-h"},
			parseResult{nil, flag.ErrHelp},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, err := parseFlags("program", tt.args)
			got := parseResult{conf, err}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf got %+v, want %+v", got, tt.want)
			}
		})
	}
}