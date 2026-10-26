package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sdrf1445/ghostOps/internal/log_handler"
)

func main() {
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	lh := log_handler.NewLogHandler(file)

	severityCount := map[log_handler.Severity]int{}
	serviceCount := map[string]int{}

	malformedLines := 0
	validLines := 0

	dbCrash := false
	dbCrashServices := map[string]bool{}
	dbCrashStartTime := time.Now()
	dbCrashEndTime := time.Time{}
	dbCrashSeverities := map[log_handler.Severity]bool{}

	index := 0

	for data, success, lastItem := lh.HandleNext(); !lastItem; data, success, lastItem = lh.HandleNext() {
		index++

		// fmt.Println(index, " ", data, " ", success)
		if !success {
			malformedLines++
			continue
		}
		validLines++
		severityCount[data.Severity]++
		serviceCount[data.Service]++

		if data.Logtype == log_handler.TypeDatabasePoolExaustion {
			dbCrash = true
			dbCrashServices[data.Service] = true
			dbCrashSeverities[data.Severity] = true
			if data.Timestamp.Before(dbCrashStartTime) {
				dbCrashStartTime = data.Timestamp
			} else if data.Timestamp.After(dbCrashEndTime) {
				dbCrashEndTime = data.Timestamp
			}
		}
	}

	fmt.Println("Malformed lines: ", malformedLines)
	fmt.Println("Valid lines: ", malformedLines)
	fmt.Println("")
	fmt.Println("Entry count by severity: ")

	for i, v := range severityCount {
		fmt.Println(log_handler.GetSeverityString(i), "", v)
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
				fmt.Println(log_handler.GetSeverityString(i))
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
