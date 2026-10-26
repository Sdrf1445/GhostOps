package main

import (
	"fmt"
	"os"

	loganalyzer "github.com/Sdrf1445/ghostOps/internal/log_analyzer"
	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

func main() {
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	// lh := loghandler.NewLogHandler(file)
	la := loganalyzer.NewLogAnalyzer(file)
	ls := la.Analyze(loghandler.TypeDatabaseConnectionTimeout | loghandler.TypeDatabasePoolExaustion | loghandler.TypeDatabasePoolThresholdApproaching | loghandler.TypePaymentGatewayTimout | loghandler.TypeDatabaseWronSyntax | loghandler.TypeDepracatedAPIEndpoint | loghandler.TypeHighMemoryUsage)

	ls.Fprint(os.Stdout)
}
