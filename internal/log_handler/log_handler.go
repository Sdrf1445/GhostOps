package loghandler

import (
	"io"
	"time"

	io_handler "github.com/Sdrf1445/ghostOps/internal/io_handler"
)

type LogHandler struct {
	iohandler *io_handler.IoHandler
}

func NewLogHandler(reader io.Reader) *LogHandler {
	ioHandler := io_handler.New(reader)

	logHandler := LogHandler{iohandler: ioHandler}

	return &logHandler
}

func (lh *LogHandler) HandleNext() (*LogData, bool, bool) {
	line, success := lh.iohandler.ReadLine()
	if !success {
		return nil, false, true
	}

	if !ReLogLine.MatchString(line) {
		return nil, false, false
	}
	match := ReLogLine.FindStringSubmatch(line)

	timestamp, err := time.Parse(time.RFC3339, match[1])
	if err != nil {
		return nil, false, false
	}

	severity := ParseSeverity(match[2])

	service := match[3]

	logType := ParseLogType(match[4])

	logdata := NewLogData(timestamp, severity, service, logType, nil)

	return logdata, true, false
}
