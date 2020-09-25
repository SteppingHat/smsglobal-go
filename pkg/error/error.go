package error

import (
	"encoding/json"
	"net/http"
)

// Error defines an error received when making a request to the API.
type Error struct {
	Code     int            `json:"code,omitempty"`
	Message  string         `json:"message,omitempty"`
	Data     []byte         `json:"data,omitempty"`
	Response *http.Response `json:"response,omitempty"` // HTTP response that caused this error
}

// Error serializes the error object and returns JSON string
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// RestAPIError represents the constructor for struct APIError
func RestAPIError(code int, message string, data []byte) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
