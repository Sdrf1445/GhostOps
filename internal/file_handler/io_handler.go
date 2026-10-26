package io_handler

import (
	"bufio"
	"io"
)

type IoHandler struct {
	reader  io.Reader
	scanner *bufio.Scanner
}

func New(reader io.Reader) *IoHandler {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	FileHandler := &IoHandler{reader: reader, scanner: scanner}
	return FileHandler
}

func (fh *IoHandler) ReadLine() (string, bool) {
	success := fh.scanner.Scan()
	if !success {
		return "", false
	}
	return fh.scanner.Text(), true
}

// scanner := bufio.NewScanner(wh.file)

// scanner.Split()
