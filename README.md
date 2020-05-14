# gfind
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
