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
Note: Special regex characters must be escaped when supplying the `-pattern` option.
