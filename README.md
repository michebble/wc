[![Go Report Card](https://goreportcard.com/badge/github.com/michebble/wc)](https://goreportcard.com/report/github.com/michebble/wc)

A cli word counter written in Go. From "Powerful Command-Line Applications in Go"

### Setup

Run tests with

```
$ go test -v
```

Build program with

```
$ go build
```

### Usage

The program accepts test via standard in, outputting the number of words via standard out.

For example

```
$ cat main.go | ./wc
151
```

The program will also accept the flag `-l` which will result in the lines being counted instead of words

```
$ cat main.go | ./wc -l
42
```
