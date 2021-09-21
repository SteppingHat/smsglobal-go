package sms

import (
	c "github.com/smsglobal/smsglobal-go/pkg/client"
	"github.com/smsglobal/smsglobal-go/util/mocks"
	"github.com/smsglobal/smsglobal-go/util/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmsGetFailed(t *testing.T) {

	client := c.New("key", "secret")

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: client,
	}
	_, err := sms.Get("6746514019161950")

	assert.Error(t, err)
}

func TestSmsGetSuccess(t *testing.T) {
	client := c.New("key", "secret")

	mocks.ResponseJson = testdata.SentToSingleDestinationResponse()
	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: client,
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

	client := c.New("key", "secret")

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	sms := &Client{
		Handler: client,
	}
	_, err := sms.List(map[string]string{})

	assert.Error(t, err)
}

func TestSmsListSuccess(t *testing.T) {
	client := c.New("key", "secret")

	mocks.ResponseJson = testdata.SmsListResponseJson()
	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	sms := &Client{
		Handler: client,
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

	client := c.New("key", "secret")
	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNotFound,
	}

	sms := &Client{
		Handler: client,
	}
	err := sms.Delete("6746514019161950")
	assert.Error(t, err)
}

func TestSmsDeleteSuccess(t *testing.T) {

	client := c.New("key", "secret")


	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNoContent,
	}

	sms := &Client{
		Handler: client,
	}

	err := sms.Delete("6746514019161950")

	if err != nil {
		t.Errorf("Sms.Delete returned error: %v", err)
	}

	assert.Nil(t, err)
}

func TestSend(t *testing.T) {

}