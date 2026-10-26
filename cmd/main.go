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

	for data, success, lastItem := lh.HandleNext(); !lastItem; data, success, lastItem = lh.HandleNext() {
		index++
		// fmt.Println(index, " ", data, " ", success)
		if !success {
			continue
		}
		severityCount[data.Severity]++
	}

	for i, v := range severityCount {
		fmt.Println(log_handler.GetSeverityString(i), "", v)
	}
}
