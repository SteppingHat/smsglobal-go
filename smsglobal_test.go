package smsglobal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/pkg/sms"
	"github.com/smsglobal/smsglobal-go/internal/pkg/sms-incoming"
	"github.com/smsglobal/smsglobal-go/internal/pkg/user"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
)

func TestNew(t *testing.T) {
	_, err := New("", "")
	assert.Error(t, err)

	s, _ := New("key", "secret")
	assert.Equal(t, "key", s.User.Handler.Key)
	assert.Equal(t, "secret", s.User.Handler.Secret)
	assert.IsType(t, s.Sms, &sms.Client{})
	assert.IsType(t, s.User, &user.Client{})
	assert.IsType(t, s.SmsIncoming, &sms_incoming.Client{})
}

func TestSMSGlobal_CreateSms(t *testing.T) {
	c, _ := New("ae311d825f75e8732b4c1a5680a11aa9", "c65fa91e71af7ccf11561e95758dd158")
	s := c.CreateSms()
	assert.IsType(t, s, &api.SendSingleSms{})
}

func TestSMSGlobal_CreateMultipleSms(t *testing.T) {
	c, _ := New("ae311d825f75e8732b4c1a5680a11aa9", "c65fa91e71af7ccf11561e95758dd158")
	s := c.CreateMultipleSms()
	assert.IsType(t, s.Messages, []*api.SendSingleSms{})
}
