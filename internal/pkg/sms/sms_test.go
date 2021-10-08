package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
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
	mocks.ResponseJson = testdata.GetSmsResponseJson()
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
	assert.Equal(t, testdata.GetSmsResponse().Id, res.Id)
	assert.Equal(t, testdata.GetSmsResponse().OutgoingId, res.OutgoingId)
	assert.Equal(t, testdata.GetSmsResponse().Origin, res.Origin)
	assert.Equal(t, testdata.GetSmsResponse().Destination, res.Destination)
	assert.Equal(t, testdata.GetSmsResponse().Message, res.Message)
	assert.Equal(t, testdata.GetSmsResponse().Status, res.Status)
	assert.Equal(t, testdata.GetSmsResponse().DateTime, res.DateTime)
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
	mocks.ResponseJson = testdata.SmsListResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	options := make(map[string]string)
	options["status"] = "undelivered"
	options["destination"] = "61401869820"
	options["startDate"] = "2020-11-23 00:00:00"

	res, err := sms.List(options)

	if err != nil {
		t.Errorf("Sms.List returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.SmsListResponse().Total, res.Total)
	assert.Equal(t, testdata.SmsListResponse().Offset, res.Offset)
	assert.Equal(t, testdata.SmsListResponse().Limit, res.Limit)
	assert.Equal(t, testdata.SmsListResponse().Messages, res.Messages)
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

func TestSmsSendOneFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	d := &api.SendSingleSms{}
	_, err := sms.SendOne(d)

	assert.Error(t, err)
}

func TestSmsSendOneSuccess(t *testing.T) {
	c := setup()

	mocks.ResponseJson = testdata.SendSmsResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	d := &api.SendSingleSms{}
	d.SetOrigin("SMSGlobal")
	d.SetDestination("61474000000")
	d.SetMessage("Message context")

	res, err := sms.SendOne(d)

	assert.Nil(t, err)
	assert.Equal(t, testdata.SendSmsResponse().Messages[0].Id, res.Messages[0].Id)
}

func TestSmsSendMultipleFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	d := &api.SendMultipleSms{}
	_, err := sms.SendMultiple(d)
	assert.Error(t, err)
}

func TestSmsSendMultipleSuccess(t *testing.T) {
	c := setup()

	mocks.ResponseJson = testdata.SendSmsResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: c,
		Logger:  l,
	}

	d := &api.SendMultipleSms{}
	d.AddMessage(&api.SendSingleSms{
		Origin:      "SMSGlobal",
		Destination: "61474000000",
		Message:     "Message context",
	})
	res, err := sms.SendMultiple(d)

	assert.Nil(t, err)
	assert.Equal(t, testdata.SendSmsResponse().Messages[0].Id, res.Messages[0].Id)
}
