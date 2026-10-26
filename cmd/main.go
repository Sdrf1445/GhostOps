package main

import (
	"fmt"
	"os"

	"github.com/Sdrf1445/ghostOps/internal/log_handler"
)

func main() {
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	lh := log_handler.NewLogHandler(file)

	index := 0
	severityCount := map[log_handler.Severity]int{}
	serviceCount := map[string]int{}

	malformedLines := 0
	validLines := 0

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
}
