package log_handler

import (
	"io"

	io_handler "github.com/Sdrf1445/ghostOps/internal/file_handler"
)

type LogHandler struct {
	iohandler *io_handler.IoHandler
}

func New(reader io.Reader) *LogHandler {
	ioHandler := io_handler.New(reader)

	logHandler := LogHandler{iohandler: ioHandler}

	return &logHandler
}

func (lh *LogHandler) HandleNext() (string, bool) {
	line, success := lh.iohandler.ReadLine()
	if !success {
		return "", false
	}
}
