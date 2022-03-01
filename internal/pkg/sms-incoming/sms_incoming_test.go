package sms_incoming

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"github.com/smsglobal/smsglobal-go/internal/util/mocks"
	"github.com/smsglobal/smsglobal-go/internal/util/testdata"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
)

var l *logger.Logger

func setup() *client.Client {
	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)
	c := client.New("key", "secret")
	c.Logger = l
	l.Debug("Setup completed")

	return c
}

func TestSmsGetFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: c,
	}
	_, err := sms.Get("6746514019161950")

	assert.Error(t, err)
}

func TestSmsGetSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.GetSmsIncomingResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	res, err := sms.Get("6746514019161950")

	if err != nil {
		t.Errorf("Sms.Get returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.GetSmsIncomingResponse(), res)

}

func TestSmsListFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}
	_, err := sms.List(map[string]string{})

	assert.Error(t, err)
}

func TestSmsListSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.SmsIncomingListResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	options := make(map[string]string)
	options["origin"] = "SMSGlobal"
	options["startDate"] = "2020-11-23 00:00:00"

	res, err := sms.List(options)

	if err != nil {
		t.Errorf("Sms.List returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.SmsIncomingListResponse().Total, res.Total)
	assert.Equal(t, testdata.SmsIncomingListResponse().Offset, res.Offset)
	assert.Equal(t, testdata.SmsIncomingListResponse().Limit, res.Limit)
	assert.Equal(t, testdata.SmsIncomingListResponse().Messages, res.Messages)
}

func TestSmsDeleteFailed(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNotFound,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}
	err := sms.Delete("6746514019161950")
	assert.Error(t, err)
}

func TestSmsDeleteSuccess(t *testing.T) {
	c := setup()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNoContent,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	err := sms.Delete("6746514019161950")

	if err != nil {
		t.Errorf("Sms.Delete returned error: %v", err)
	}

	assert.Nil(t, err)
}
