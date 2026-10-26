package log_handler

import "time"

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

type LogType int

const (
	TypeUnkown LogType = iota
	TypeUserLogin
	TypeTokenGeneration
	TypeDatabasePoolExaustion
	TypeDatabasePoolThresholdApproaching
	TypeDatabaseConnectionTimeout
	TypeDatabaseWronSyntax
	TypeHighMemoryUsage
	TypeDepracatedAPIEndpoint
	TypePaymentGatewayTimout
	TypeHttp
	TypeRateLimitApproaching
)

type HttpMethod int

const (
	MethodGet HttpMethod = iota
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
)

type TypeHttpData struct {
	method      string
	path        string
	status      int
	duration_ms int
}

type TypeRateLimitData struct {
	clientip string
}

type TypeDepractedEndpointData struct {
	endpoint string
}

type TypeDatabasePoolThresholdData struct {
	util_percent int
}

type TypeHighMemoryUsageData struct {
	memory_percent int
}

type TypePaymentGatewayTimoutData struct {
	status int
	reason string
}

type LogData struct {
	timestamp time.Time
	severity  Severity
	service   string
	data      any
}
