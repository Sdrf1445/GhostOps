package logsummary

import (
	"time"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogTypeData struct {
	LogTypePresent    bool
	LogTypeServices   map[string]bool
	LogTypeStartTime  time.Time
	LogTypeEndTime    time.Time
	LogTypeSeverities map[loghandler.Severity]bool
}

func NewLogTypeData() *LogTypeData {
	logTypeServices := map[string]bool{}
	logTypeStartTime := time.Now()
	logTypeEndTime := time.Time{}
	logTypeSeverities := map[loghandler.Severity]bool{}

	ltd := LogTypeData{LogTypeServices: logTypeServices, LogTypeStartTime: logTypeStartTime, LogTypeEndTime: logTypeEndTime, LogTypeSeverities: logTypeSeverities, LogTypePresent: false}

	return &ltd
}
