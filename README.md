# Rqueue

[![Build Status](https://travis-ci.org/wellmart/rqueue.svg?branch=master)](https://travis-ci.org/wellmart/rqueue)
[![Go Report Card](https://goreportcard.com/badge/github.com/wellmart/rqueue)](https://goreportcard.com/report/github.com/wellmart/rqueue)
[![Coverage Status](https://coveralls.io/repos/github/wellmart/rqueue/badge.svg?branch=master)](https://coveralls.io/github/wellmart/rqueue?branch=master)
![Version](https://img.shields.io/badge/version-0.1.0-blue)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wellmart/rqueue?status.svg)](https://godoc.org/github.com/wellmart/rqueue)

A queue is a collection of items following a first-in, first-out principle.

## Requirements

Go 1.1 and beyond.

## Installation

Use the go package manager to install rqueue.

```bash
go get github.com/wellmart/rqueue
```

## Usage

```go
package main


import (
	"github.com/wellmart/rqueue"
)

func main() {
	queue := New()

	queue.Enqueue(1)
	queue.TryDequeue()
}
```

## Staying up to date

To update rqueue to the latest version, use `go get -u github.com/wellmart/rqueue`.

## License

[MIT](https://choosealicense.com/licenses/mit/)
