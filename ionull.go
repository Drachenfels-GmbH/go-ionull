package main

import (
	"io"
)

type null struct{}

var Null = null{}

// Consumes and discards everything.
func (n null) Write(p []byte) (int, error) {
	return len(p), nil
}

// Always returns a size of 0 and io.EOF
func (n null) Read(p []byte) (int, error) {
	return 0, io.EOF
}
