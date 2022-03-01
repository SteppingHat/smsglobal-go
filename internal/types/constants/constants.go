package constants

const (
	Host        = "https://api.smsglobal.com/v2"
	ContentType = "application/json"
	Version     = "1.0.0"
	DefaultCode = -1 //defaultCode is the default error code for non-api related failures. e.g. Missing credentials
	DebugLevel  = "debug"
	Timeout     = 30 //request timeout duration in seconds
	DateTimeFormat = "2006-01-02 15:04:05" // The datetime format expected by rest api  yyyy-MM-dd HH:mm:ss
)
