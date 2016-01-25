# Iggy

![baby iggy](https://cloud.githubusercontent.com/assets/836375/9753627/13b8bc50-5685-11e5-940d-dc4c6c3473cb.png)

Grabs a gitignore for you, and writes it to `.gitignore`.

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

## Building and Contributing

This package does fun things like generating go code *for the developer*. This is primarily done with a single shell script, `gen.sh`. To get the show on the road, we'll be using `go generate` followed by `go build` (which creates `iggy`:

To build from a fresh clone of this repository:

```
go generate
go build
```

The `go generate` triggers a shell script called `gen.sh` which clones github.com/github/gitignore and writes a big `map[string]string` to `langs/langs.go` of all the gitignores.

Hacking on this package is more about that shell script and the cli interface. Pull requests, patches, and general praise are well appreciated. :smile:

## Roadmap

* [X] Generate gitignores based on name
* [ ] Allow for aliases (e.g. node, Node, iojs, JavaScript)
* [ ] Customize codegangsta/cli to make simpler output
* [ ] Provide option of writing to stdout instead of file
* [ ] Detect language (and prompt) when language not provided
* [ ] Create built binaries *for all*!
