package logsummary

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogSummary struct {
	malformedLines int
	validLines     int
	severities     map[loghandler.Severity]int
	services       map[string]int
	incidents      map[loghandler.LogType]*LogTypeData
}

func New(malformedLines int, validLines int, severities map[loghandler.Severity]int, services map[string]int, incidents map[loghandler.LogType]*LogTypeData) *LogSummary {
	ls := LogSummary{malformedLines: malformedLines, validLines: validLines, severities: severities, services: services, incidents: incidents}
	return &ls
}

func (ls *LogSummary) Fprint(writer io.Writer) {
	fmt.Fprintf(writer, "Malformed lines: %d\n", ls.malformedLines)
	fmt.Fprintf(writer, "Valid lines: %d\n", ls.validLines)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Severity count:\n")
	for severity, count := range ls.severities {
		fmt.Fprintf(writer, "%s %d\n", loghandler.GetSeverityString(severity), count)
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Service count:\n")
	for service, count := range ls.services {
		fmt.Fprintf(writer, "%s %d\n", service, count)
	}
	fmt.Fprintln(writer)

	for incident, incidentData := range ls.incidents {
		if !incidentData.LogTypePresent {
			continue
		}
		fmt.Fprintf(writer, "%s:\n", incident.GetLogTypeString())
		fmt.Fprintf(writer, "Severities:\n")
		for severity, exists := range incidentData.LogTypeSeverities {
			if exists {
				fmt.Fprintf(writer, "%s\n", loghandler.GetSeverityString(severity))
			}
		}
		fmt.Fprintf(writer, "Services:\n")
		for service, exists := range incidentData.LogTypeServices {
			if exists {
				fmt.Fprintf(writer, "%s\n", service)
			}
		}
		fmt.Fprint(writer, "Time start: ", incidentData.LogTypeStartTime, "\n")
		fmt.Fprint(writer, "Time end: ", incidentData.LogTypeEndTime, "\n")
		fmt.Fprintln(writer)
	}
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
