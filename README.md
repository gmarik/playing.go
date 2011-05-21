## My Go lanuage experiments

Mostly notes for myself

## Getting started with Go on OSX

    $ brew update

    $ brew install go
    Error: You must `easy_install mercurial'

    $ easy_install mercurial
    Searching for mercurial
    ...
    Installing hg script to /usr/local/bin

    $ brew install go
    /usr/local//bin/hg
    ==> Cloning http://go.googlecode.com/hg/
    Updating /Users/marjanhratson/Library/Caches/Homebrew/go--hg
    pulling from http://go.googlecode.com/hg/
    searching for changes
    no changes found
    0 files updated, 0 files merged, 0 files removed, 0 files unresolved
    ==> Checking out revision release
    ==> ./all.bash
    /usr/local/Cellar/go/HEAD: 1797 files, 104M, built in 4.7 minutes

  and [RTFM](http://golang.org/doc/docs.html)

## Hello world with Go

see [hello world example](https://github.com/gmarik/playing.go/blob/master/helloworld/helloworld.go)

## Writing Test for Go

- read [How to Write Go Code](http://golang.org/doc/code.html#pkg_example)

## More Go examples

- [gmallard/go-samp](https://github.com/gmallard/go-samp)
- [tokuhirom/go-examples](https://github.com/tokuhirom/go-examples)
