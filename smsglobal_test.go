package smsglobal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	_, err := New("", "")
	assert.Error(t, err)

	s, _ := New("key", "secret")
	assert.Equal(t, s.User.Handler.Key, "key")
	assert.Equal(t, s.User.Handler.Secret, "secret")
}
