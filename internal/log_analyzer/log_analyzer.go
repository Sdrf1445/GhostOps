package loganalyzer

import (
	"io"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
	logsummary "github.com/Sdrf1445/ghostOps/internal/log_summary"
)

type LogAnalyzer struct {
	logHandler *loghandler.LogHandler
}

func NewLogAnalyzer(reader io.Reader) *LogAnalyzer {
	lh := loghandler.NewLogHandler(reader)
	la := LogAnalyzer{logHandler: lh}
	return &la
}

func (la *LogAnalyzer) Analyze(logTypes loghandler.LogType) *logsummary.LogSummary {
	severityCount := map[loghandler.Severity]int{}
	serviceCount := map[string]int{}

	malformedLines := 0
	validLines := 0

	logTypeMap := map[loghandler.LogType]*logsummary.LogTypeData{}

	resultLogTypes := logTypes.GetAllLogTypes()

	for _, v := range resultLogTypes {
		ltd := logsummary.NewLogTypeData()
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
				logTypeMap[v].LogTypePresent = true
				logTypeMap[v].LogTypeServices[data.Service] = true
				logTypeMap[v].LogTypeSeverities[data.Severity] = true
				if data.Timestamp.Before(logTypeMap[v].LogTypeStartTime) {
					logTypeMap[v].LogTypeStartTime = data.Timestamp
				} else if data.Timestamp.After(logTypeMap[v].LogTypeEndTime) {
					logTypeMap[v].LogTypeEndTime = data.Timestamp
				}
			}
		}
	}

	ls := logsummary.New(malformedLines, validLines, severityCount, serviceCount, logTypeMap)
	return ls

	// fmt.Println("Malformed lines: ", malformedLines)
	// fmt.Println("Valid lines: ", malformedLines)
	// fmt.Println("")
	// fmt.Println("Entry count by severity: ")

	// for i, v := range severityCount {
	// 	fmt.Println(loghandler.GetSeverityString(i), "", v)
	// }
	// fmt.Println("")
	// fmt.Println("Entry count by service: ")
	// for i, v := range serviceCount {
	// 	fmt.Println(i, "", v)
	// }
	// fmt.Println("")
	// if dbCrash {
	// 	fmt.Println("Database connection exaustion crash")
	// 	fmt.Println("Severities:")
	// 	for i, v := range dbCrashSeverities {
	// 		if v {
	// 			fmt.Println(loghandler.GetSeverityString(i))
	// 		}
	// 	}
	// 	fmt.Println("Services:")
	// 	for i, v := range dbCrashServices {
	// 		if v {
	// 			fmt.Println(i)
	// 		}
	// 	}
	// 	fmt.Println("Time start:", dbCrashStartTime)
	// 	fmt.Println("Time end:", dbCrashEndTime)
	// }
}
