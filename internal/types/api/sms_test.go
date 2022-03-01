package api

import (
	"testing"
	"time"

	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"github.com/stretchr/testify/assert"
)

func TestSendSingleSms_SetOrigin(t *testing.T) {
	s := &SendSingleSms{}
	s.SetOrigin("SMSGlobal")
	assert.Equal(t, "SMSGlobal", s.Origin)
}

func TestSendSingleSms_SetDestination(t *testing.T) {
	s := &SendSingleSms{}
	s.SetDestination("61474000000")
	assert.Equal(t, "61474000000", s.Destination)
}

func TestSendSingleSms_AddDestination(t *testing.T) {
	s := &SendSingleSms{}
	s.AddDestination("61474000000")
	s.AddDestination("61474950850")

	assert.Contains(t, s.Destinations, "61474950850")
}

func TestSendSingleSms_SetMessage(t *testing.T) {
	s := &SendSingleSms{}
	s.SetMessage("Message content")

	assert.NotEmpty(t, s.Message)
	assert.Equal(t, "Message content", s.Message)
}

func TestSendSingleSms_SetScheduledDateTime(t *testing.T) {
	s := &SendSingleSms{}
	now := time.Now()
	s.SetScheduledDateTime(now)
	assert.NotEmpty(t, s.ScheduledDateTime)
	assert.Equal(t, now.Format(constants.DateTimeFormat), s.ScheduledDateTime)
}

func TestSendSingleSms_SetCampaign(t *testing.T) {
	s := &SendSingleSms{}
	s.SetCampaign("1234")

	assert.NotEmpty(t, s.Campaign)
	assert.Equal(t, "1234", s.Campaign)
}

func TestSendSingleSms_SetSharedPool(t *testing.T) {
	s := &SendSingleSms{}
	s.SetSharedPool("1234")

	assert.NotEmpty(t, s.SharedPool)
	assert.Equal(t, "1234", s.SharedPool)
}

func TestSendSingleSms_SetNotifyUrl(t *testing.T) {
	s := &SendSingleSms{}
	s.SetNotifyUrl("https://notification.callback.com")

	assert.NotEmpty(t, s.NotifyUrl)
	assert.Equal(t, "https://notification.callback.com", s.NotifyUrl)
}

func TestSendSingleSms_SetIncomingUrl(t *testing.T) {
	s := &SendSingleSms{}
	s.SetIncomingUrl("https://incoming.message.com")

	assert.NotEmpty(t, s.IncomingUrl)
	assert.Equal(t, "https://incoming.message.com", s.IncomingUrl)
}

func TestSendSingleSms_SetExpiryDateTime(t *testing.T) {
	s := &SendSingleSms{}
	now := time.Now()
	s.SetExpiryDateTime(now)
	assert.NotEmpty(t, s.ExpiryDateTime)
	assert.Equal(t, s.ExpiryDateTime, now.Format(constants.DateTimeFormat))
}

func TestSendMultipleSms_AddMessage(t *testing.T) {
	s := &SendSingleSms{}

	s.SetMessage("Message content")
	s.SetOrigin("SMSGlobal")
	s.SetDestination("61474000000")
	e := time.Now().Add(time.Hour * 10)
	s.SetExpiryDateTime(e)

	m := &SendMultipleSms{}
	m.AddMessage(s)

	assert.Equal(t, m.Messages, []*SendSingleSms{
		{
			Origin:         "SMSGlobal",
			Destination:    "61474000000",
			Message:        "Message content",
			ExpiryDateTime: e.Format(constants.DateTimeFormat),
		},
	})
}
