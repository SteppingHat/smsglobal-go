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
