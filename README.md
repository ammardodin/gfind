# gfind

[![CircleCI](https://circleci.com/gh/ammardodin/gfind.svg?style=shield)](https://circleci.com/gh/ammardodin/gfind)

Concurrent file finder in golang.

## How to Run ?
1. Clone the repository.
2. Set your `GOPATH` to the repository's root.
3. Run `make install` in the repository's root.
4. Run gfind like this:
```shell script
gfind -start <path-to-starting-directory> -pattern <pattern-to-match-against>
```

## Notes
1. Special regex characters must be escaped when supplying the `-pattern` option.
2. By default, much like GNU's `find`, `gfind` will output all errors it encounters to `stderr`. If you would like to silence said errors, make sure to redirect `stderr` to the `null` device, i.e. tack on a `2>/dev/null` to the end of the command.

## Example
Suppose the absolute path to the current working directory is `/Users/batman/awesomeProject` and it has a hierarchy as follows:
```shell script
.
├── Makefile
├── README.md
├── bin
├── pkg
└── src
    └── user
        └── gfind
            ├── find
            │   └── versions
            │       ├── v1
            │       │   └── finder_v1.go
            │       └── v2
            │           └── finder_v2.go
            ├── finder.go
            ├── finder_test.go
            ├── main.go
            ├── parse_flags_test.go
            ├── string_queue.go
            └── string_queue_test.go
```
If we were to run:
```shell script
gfind -start . -pattern "finder.*\.go" 2>/dev/null
```
The output would be:
```shell script
start: /Users/batman/awesomeProject, pattern: finder.*\.go
Using 12 workers !
/Users/batman/awesomeProject/src/user/gfind/finder.go
/Users/batman/awesomeProject/src/user/gfind/finder_test.go
/Users/batman/awesomeProject/src/user/gfind/find/versions/v1/finder_v1.go
/Users/batman/awesomeProject/src/user/gfind/find/versions/v2/finder_v2.go
```
