package logsummary

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	loghandler "github.com/Sdrf1445/ghostOps/internal/log_handler"
)

type LogSummary struct {
	MalformedLines int
	ValidLines     int
	Severities     map[loghandler.Severity]int
	Services       map[string]int
	Incidents      map[loghandler.LogType]*LogTypeData
}

func New(malformedLines int, validLines int, severities map[loghandler.Severity]int, services map[string]int, incidents map[loghandler.LogType]*LogTypeData) *LogSummary {
	ls := LogSummary{MalformedLines: malformedLines, ValidLines: validLines, Severities: severities, Services: services, Incidents: incidents}
	return &ls
}

func (ls *LogSummary) Fprint(writer io.Writer) {
	fmt.Fprintf(writer, "Malformed lines: %d\n", ls.MalformedLines)
	fmt.Fprintf(writer, "Valid lines: %d\n", ls.ValidLines)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Severity count:\n")
	for severity, count := range ls.Severities {
		fmt.Fprintf(writer, "%s %d\n", loghandler.GetSeverityString(severity), count)
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Service count:\n")
	for service, count := range ls.Services {
		fmt.Fprintf(writer, "%s %d\n", service, count)
	}
	fmt.Fprintln(writer)

	for incident, incidentData := range ls.Incidents {
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
