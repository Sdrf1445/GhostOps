package logsummary

import (
	"encoding/json"
	"os"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogSummary struct {
	malformedLines int
	validLines     int
	severities     map[loghandler.Severity]int
	services       map[string]int
	// incidents      map[loghandler.LogType]loganalyzer.LogTypeData
}

func New(malformedLines int, validLines int, severities map[loghandler.Severity]int, services map[string]int) *LogSummary {
	ls := LogSummary{malformedLines: malformedLines, validLines: validLines, severities: severities, services: services}
	return &ls
}

func (ls *LogSummary) Save(filePath string) error {
	data, err := json.MarshalIndent(ls, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}
