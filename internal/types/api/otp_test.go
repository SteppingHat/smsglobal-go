package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/types/constants"
)

func TestSendOtp_SetOrigin(t *testing.T) {
	o := &SendOtp{}

	o.SetOrigin("SMSGlobal")

	assert.Equal(t, "SMSGlobal", o.Origin)
}

func TestSendOtp_SetDestination(t *testing.T) {
	o := &SendOtp{}
	o.SetDestination("61474000000")
	assert.Equal(t, "61474000000", o.Destination)
}

func TestSendOtp_SetLength(t *testing.T) {
	o := &SendOtp{}
	o.SetLength(8)
	assert.Equal(t, 8, o.Length)
}

func TestSendOtp_SetMessage(t *testing.T) {
	o := &SendOtp{}
	o.SetMessage("{*code*} is your SMSGlobal verification code.")
	assert.Equal(t, "{*code*} is your SMSGlobal verification code.", o.Message)
}

func TestSendOtp_SetCodeExpiry(t *testing.T) {
	o := &SendOtp{}
	o.SetCodeExpiry(600)
	assert.Equal(t, 600, o.CodeExpiry)
}

func TestSendOtp_SetMessageExpiryDateTime(t *testing.T) {
	o := &SendOtp{}

	e := time.Now().Add(time.Minute * 30)
	o.SetMessageExpiryDateTime(e)

	assert.Equal(t, e.Format(constants.DateTimeFormat), o.MessageExpiryDateTime)
}
