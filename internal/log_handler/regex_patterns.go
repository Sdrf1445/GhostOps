package log_handler

import "regexp"

var (
	ReLogLine     = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z) (INFO|WARN|FATAL|ERROR) (\S+) (.+?)`)
	ReHttp        = regexp.MustCompile(`^HTTP ([A-Z]+) (\S+) status=([0-9]{3}) duration=([0-9]+)ms$`)
	ReRateLimit   = regexp.MustCompile(`^Rate limit approaching for client IP (\S+)$`)
	ReDeprecated  = regexp.MustCompile(`^Deprecated API endpoint (\S+) accessed$`)
	ReDbThreshold = regexp.MustCompile(`^Database connection pool reaching threshold \(([0-9]+)% utilized\)$`)
	ReMemory      = regexp.MustCompile(`^High memory usage detected: ([0-9]+)%$`)
	ReDbTimeout   = regexp.MustCompile(`^Failed to connect to database: timeout after ([0-9]+)ms$`)
	ReDateway502  = regexp.MustCompile(`^Payment gateway returned ([0-9]{3}) (.+)$`)
	// npeRe           = regexp.MustCompile(`^NullPointerException in controller layer processing request$`)
	ReSqlSyntax       = regexp.MustCompile(`^SQL syntax error while executing query: (.+)$`)
	RePayTimeout      = regexp.MustCompile(`^Payment gateway timeout - transaction aborted$`)
	ReDbExhaust       = regexp.MustCompile(`^Database connection pool exhausted and crashed system$`)
	ReLoginSuccess    = regexp.MustCompile(`^User successfully logged in$`)
	ReTokenGenerated  = regexp.MustCompile(`^Token generated for user session$`)
	ReCacheRefreshed  = regexp.MustCompile(`^Cache refreshed successfully$`)
	ReWebhookReceived = regexp.MustCompile(`^Incoming webhook received from gateway$`)
)
