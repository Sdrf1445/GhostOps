package main

import (
	"fmt"
	"os"

	loganalyzer "github.com/Sdrf1445/ghostOps/internal/log_analyzer"
	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 4 {
		fmt.Printf("Usage:go run ./cmd/main.go <log file path> [-o <output json file>]")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var outputFile string
	if len(os.Args) > 2 && os.Args[2] == "-o" {
		outputFile = os.Args[3]
	}

	la := loganalyzer.NewLogAnalyzer(file)
	ls := la.Analyze(loghandler.TypeDatabaseConnectionTimeout | loghandler.TypeDatabasePoolExaustion | loghandler.TypeDatabasePoolThresholdApproaching | loghandler.TypePaymentGatewayTimout | loghandler.TypeDatabaseWronSyntax | loghandler.TypeDepracatedAPIEndpoint | loghandler.TypeHighMemoryUsage)

	ls.Fprint(os.Stdout)

	if outputFile != "" {
		err := ls.Save(outputFile)
		if err != nil {
			fmt.Println("Could not save the file: ", err)
			return
		}
		fmt.Println("File saved successfuly")
	}
}
