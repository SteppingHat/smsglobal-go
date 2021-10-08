package error

import (
	"encoding/json"
)

// Error defines an error received when making a request to the API.
type Error struct {
	Code     int            `json:"code,omitempty"`
	Message  string         `json:"message,omitempty"`
	Data     []byte         `json:"data,omitempty"`
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

