package loganalyzer

import (
	"fmt"
	"io"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogAnalyzer struct {
	logHandler *loghandler.LogHandler
}

func NewLogAnalyzer(reader io.Reader) *LogAnalyzer {
	lh := loghandler.NewLogHandler(reader)
	la := LogAnalyzer{logHandler: lh}
	return &la
}

func (la *LogAnalyzer) Analyze(logTypes loghandler.LogType) {
	severityCount := map[loghandler.Severity]int{}
	serviceCount := map[string]int{}

	malformedLines := 0
	validLines := 0

	logTypeMap := map[loghandler.LogType]*LogTypeData{}

	resultLogTypes := logTypes.GetAllLogTypes()

	for _, v := range resultLogTypes {
		ltd := NewLogTypeData()
		logTypeMap[v] = ltd
	}

	// dbCrash := false
	// dbCrashServices := map[string]bool{}
	// dbCrashStartTime := time.Now()
	// dbCrashEndTime := time.Time{}
	// dbCrashSeverities := map[loghandler.Severity]bool{}

	index := 0

	for data, success, lastItem := la.logHandler.HandleNext(); !lastItem; data, success, lastItem = la.logHandler.HandleNext() {
		index++

		// fmt.Println(index, " ", data, " ", success)
		if !success {
			malformedLines++
			continue
		}
		validLines++
		severityCount[data.Severity]++
		serviceCount[data.Service]++

		for _, v := range resultLogTypes {
			if data.Logtype == v {
				logTypeMap[v].logTypePresent = true
				logTypeMap[v].logTypeServices[data.Service] = true
				logTypeMap[v].logTypeSeverities[data.Severity] = true
				if data.Timestamp.Before(logTypeMap[v].logTypeStartTime) {
					logTypeMap[v].logTypeStartTime = data.Timestamp
				} else if data.Timestamp.After(logTypeMap[v].logTypeEndTime) {
					logTypeMap[v].logTypeEndTime = data.Timestamp
				}
			}
		}
	}

	fmt.Println("Malformed lines: ", malformedLines)
	fmt.Println("Valid lines: ", malformedLines)
	fmt.Println("")
	fmt.Println("Entry count by severity: ")

	for i, v := range severityCount {
		fmt.Println(loghandler.GetSeverityString(i), "", v)
	}
	fmt.Println("")
	fmt.Println("Entry count by service: ")
	for i, v := range serviceCount {
		fmt.Println(i, "", v)
	}
	fmt.Println("")
	if dbCrash {
		fmt.Println("Database connection exaustion crash")
		fmt.Println("Severities:")
		for i, v := range dbCrashSeverities {
			if v {
				fmt.Println(loghandler.GetSeverityString(i))
			}
		}
		fmt.Println("Services:")
		for i, v := range dbCrashServices {
			if v {
				fmt.Println(i)
			}
		}
		fmt.Println("Time start:", dbCrashStartTime)
		fmt.Println("Time end:", dbCrashEndTime)
	}
}
