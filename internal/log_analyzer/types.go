package loganalyzer

import (
	"time"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogTypeData struct {
	logTypePresent    bool
	logTypeServices   map[string]bool
	logTypeStartTime  time.Time
	logTypeEndTime    time.Time
	logTypeSeverities map[loghandler.Severity]bool
}

func NewLogTypeData() *LogTypeData {
	logTypeServices := map[string]bool{}
	logTypeStartTime := time.Now()
	logTypeEndTime := time.Time{}
	logTypeSeverities := map[loghandler.Severity]bool{}

	ltd := LogTypeData{logTypeServices: logTypeServices, logTypeStartTime: logTypeStartTime, logTypeEndTime: logTypeEndTime, logTypeSeverities: logTypeSeverities, logTypePresent: false}

	return &ltd
}
