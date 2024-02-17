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

The program accepts test via standard input

By default it will count the number of words.

```
$ echo "This is sample text\nFollowed by more text on a new line" | ./wc
4
```

Pass the flag `-l` to count the number of lines instead of words.

```
$ echo "This is sample text\nFollowed by more text on a new line" | ./wc -l
2
```

Pass the flag `-b` to count the number of bytes instead.

```
$ echo "This is sample text\nFollowed by more text on a new line" | ./wc -b
56
```
