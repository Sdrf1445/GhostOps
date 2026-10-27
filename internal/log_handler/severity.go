package loghandler

type Severity int

const (
	SevUnknown Severity = iota
	SevINFO
	SevWARN
	SevERROR
	SevFATAL
)

func ParseSeverity(severity string) Severity {
	switch severity {
	case "INFO":
		return SevINFO
	case "WARN":
		return SevWARN
	case "ERROR":
		return SevERROR
	case "FATAL":
		return SevFATAL
	default:
		return SevUnknown
	}
}

func GetSeverityString(severity Severity) string {
	switch severity {
	case SevINFO:
		return "INFO"
	case SevWARN:
		return "WARN"
	case SevERROR:
		return "ERROR"
	case SevFATAL:
		return "FATAL"
	default:
		return "Unkown"
	}
}

func (s Severity) MarshalText() ([]byte, error) {
	return []byte(GetSeverityString(s)), nil
}
