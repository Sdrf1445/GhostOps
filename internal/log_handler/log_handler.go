package log_handler

import (
	"io"
	"regexp"

	io_handler "github.com/Sdrf1445/ghostOps/internal/file_handler"
)

type LogHandler struct {
	iohandler *io_handler.IoHandler
}

func NewLogHandler(reader io.Reader) *LogHandler {
	ioHandler := io_handler.New(reader)

	logHandler := LogHandler{iohandler: ioHandler}

	return &logHandler
}

var re = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z) (INFO|WARN|FATAL|ERROR) (\S+) (.+?)`)

func (lh *LogHandler) HandleNext() (string, bool) {
	line, success := lh.iohandler.ReadLine()
	if !success {
		return "", false
	}

	if !re.MatchString(line) {
		return "", false
	}
}
