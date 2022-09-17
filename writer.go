package log

import "io"

type Writer interface {
	io.Writer
	Sync() error
	Flush() error
}

