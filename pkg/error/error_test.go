package error

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := Error{Code: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest)}
	assert.Equal(t, `{"code":400,"message":"Bad Request"}`, err.Error())
}

// Test api error constructor
func TestRestAPIError(t *testing.T) {

	var data []byte
	err := RestAPIError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), data)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, err.Code)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), err.Message)
}