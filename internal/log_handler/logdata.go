package loghandler

import "time"

type LogData struct {
	Timestamp time.Time
	Severity  Severity
	Service   string
	Logtype   LogType
	Data      any
}

func NewLogData(timestamp time.Time, severity Severity, service string, logtype LogType, data any) *LogData {
	logData := LogData{Timestamp: timestamp, Severity: severity, Service: service, Data: data, Logtype: logtype}
	return &logData
}
