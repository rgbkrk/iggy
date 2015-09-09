# Iggy

Grabs gitignores for you, writes them to `.gitignore`.

## Why

I like creating repos from the command line and adding a gitignore right after. GitHub has spoiled me with autocreated repos. I figured it was about time I wrote a simple binary tool that would dump out a gitignore for a given language.

## Usage

```
$ iggy Go
```

writes out a `.gitignore` consisting of

```
# Compiled Object files, Static and Dynamic libs (Shared Objects)
*.o
*.a
*.so

# Folders
_obj
_test

# Architecture specific extensions/prefixes
*.[568vq]
[568vq].out

*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*

_testmain.go

*.exe
*.test
*.prof
```

The same can be done for Node, Python, whatever is supported in [github/gitignore](https://github.com/github/gitignore).

## Building

This package does fun things like generating go code *for the developer*. This is primarily done with a single shell script, `gen.sh`. To get the show on the road, we'll be using `go generate` followed by `go build` (which creates `iggy`:

```
go generate
go build
```

## Roadmap

* [X] Generate gitignores based on name
* [ ] Allow for aliases (e.g. node, Node, iojs, JavaScript)
* [ ] Customize codegangsta/cli to make simpler output
* [ ] Provide option of writing to stdout instead of file
* [ ] Detect language (and prompt) when language not provided
* [ ] Create built binaries *for all*!
