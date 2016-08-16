package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io"
)

func TestRead(t *testing.T) {
	p := make([]byte, 20)
	n, err := Null.Read(p)
	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)
}

func TestWrite(t *testing.T) {
	p := []byte("foobar")
	n, err := Null.Write(p)
	assert.Equal(t, len(p), n)
	assert.NoError(t, err)
}
