package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSingleSms_SetOrigin(t *testing.T) {
	s := &SendSingleSms{}
	s.SetOrigin("SMSGlobal")
	assert.Equal(t, s.Origin, "SMSGlobal")
}

func TestSendSingleSms_SetDestination(t *testing.T) {
	s := &SendSingleSms{}
	s.SetDestination("61474000000")
	assert.Equal(t, s.Destination, "61474000000")
}

func TestSendSingleSms_AddDestination(t *testing.T) {
	s := &SendSingleSms{}
	s.AddDestination("61474000000")
	s.AddDestination("61474950850")

	assert.Contains(t, s.Destinations, "61474950850")

}

func TestSendSingleSms_SetMessage(t *testing.T) {
	s := &SendSingleSms{}
	s.SetMessage("Message context")

	assert.NotEmpty(t, s.Message)
	assert.Equal(t, s.Message, "Message context")
}

func TestSendMultipleSms_AddMessage(t *testing.T) {
	s := &SendMultipleSms{}

	s.AddMessage(&SendSingleSms{
		Origin:      "SMSGlobal",
		Destination: "61474000000",
		Message:     "Message context",
	})

	assert.ObjectsAreEqual(s.Messages, []*SendSingleSms{
		{
			Origin:      "SMSGlobal",
			Destination: "61474000000",
			Message:     "Message context",
		},
	})
}
