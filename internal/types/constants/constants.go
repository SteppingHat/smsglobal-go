package constants

const (
	Host        = "https://api.smsglobal.com/v2"
	ContentType = "application/json"
	Version     = "1.0.0"
	UserAgent   = "SMSGlobal-GO-SDK/" + Version
	DefaultCode = -1 //defaultCode is the default error code for non-api related failures. eg. Missing credentials
	DebugLevel  = "info"
	Timeout     = 30 //request timeout duration in seconds
)
