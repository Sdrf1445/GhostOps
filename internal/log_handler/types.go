package loghandler

import (
	"time"
)

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

type LogType int

const (
	TypeUnkown LogType = 1 << iota
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
	TypeCacheRefresh
	TypeWebhookRecieved
)

func (logtype LogType) HasFlag(input LogType) bool {
	return logtype&input != 0
}

func (logTypes LogType) GetAllLogTypes() []LogType {
	result := []LogType{}
	for logTypes != 0 {
		flag := logTypes & -logTypes

		result = append(result, flag)

		logTypes &= logTypes - 1
	}
	return result
}

func (logtype *LogType) GetLogTypeString() string {
	switch *logtype {
	case TypeUnkown:
		return "Unknown"
	case TypeUserLogin:
		return "User logins"
	case TypeTokenGeneration:
		return "User token generations"
	case TypeDatabasePoolExaustion:
		return "Database connection pool exhaustion"
	case TypeDatabasePoolThresholdApproaching:
		return "Database connection pool threshold approaching"
	case TypeDatabaseConnectionTimeout:
		return "Database connection timeout"
	case TypeDatabaseWronSyntax:
		return "Database wrong syntax"
	case TypeHighMemoryUsage:
		return "High memory usage"
	case TypeDepracatedAPIEndpoint:
		return "Deprecated API endpoint"
	case TypePaymentGatewayTimout:
		return "Payment gateway timeout"
	case TypeHttp:
		return "HTTP"
	case TypeRateLimitApproaching:
		return "Rate limit approaching"
	case TypeCacheRefresh:
		return "Cache refresh"
	case TypeWebhookRecieved:
		return "Webhook received"
	default:
		return "Unknown log type"
	}
}

func ParseLogType(message string) LogType {
	if ReLoginSuccess.MatchString(message) {
		return TypeUserLogin
	}
	if ReTokenGenerated.MatchString(message) {
		return TypeTokenGeneration
	}
	if ReDbExhaust.MatchString(message) {
		return TypeDatabasePoolExaustion
	}
	if ReDbThreshold.MatchString(message) {
		return TypeDatabasePoolThresholdApproaching
	}
	if ReDbTimeout.MatchString(message) {
		return TypeDatabaseConnectionTimeout
	}
	if ReSqlSyntax.MatchString(message) {
		return TypeDatabaseWronSyntax
	}
	if ReMemory.MatchString(message) {
		return TypeHighMemoryUsage
	}
	if ReDeprecated.MatchString(message) {
		return TypeDepracatedAPIEndpoint
	}
	if RePayTimeout.MatchString(message) {
		return TypePaymentGatewayTimout
	}
	if ReHttp.MatchString(message) {
		return TypeHttp
	}
	if ReRateLimit.MatchString(message) {
		return TypeRateLimitApproaching
	}
	if ReCacheRefreshed.MatchString(message) {
		return TypeCacheRefresh
	}
	if ReWebhookReceived.MatchString(message) {
		return TypeWebhookRecieved
	}
	return TypeUnkown
}

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
	Timestamp time.Time
	Severity  Severity
	Service   string
	Logtype   LogType
	Data      any
}

func NewLogData(timestamp time.Time, severity Severity, service string, logtype LogType, data any) *LogData {
	logData := LogData{Timestamp: timestamp, Severity: severity, Service: service, Data: data, Logtype: logtype}
	return &logData
}
