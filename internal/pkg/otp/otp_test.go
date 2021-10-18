package otp

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

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

func TestSendOtpFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	d := &api.SendOtp{}
	_, err := otp.Send(d)

	assert.Error(t, err)
}

func TestSendOtpSuccess(t *testing.T) {
	c := setup()

	mocks.ResponseJson = testdata.SendOtpResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	o := &api.SendOtp{}

	o.SetOrigin("SMSGlobal")
	o.SetDestination("61474000000")
	o.SetMessage("{*code*} is your SMSGlobal verification code.")
	o.SetCodeExpiry(600)
	o.SetLength(4)

	e := time.Now().Add(time.Minute * 30)
	o.SetMessageExpiryDateTime(e)
	res, err := otp.Send(o)

	assert.Nil(t, err)
	assert.ObjectsAreEqual(testdata.SendOtpResponse(), res)
	assert.Equal(t, "Sent", testdata.SendOtpResponse().Status)
}

func TestGenerateMsisdnPath(t *testing.T) {
	assert.Equal(t, "/otp/61400000000/validate", generateMsisdnPath("61400000000", "validate"))
	assert.Equal(t, "/otp/61400000000/cancel", generateMsisdnPath("61400000000", "cancel"))
}

func TestGenerateRequestIdPath(t *testing.T) {
	assert.Equal(t, "/otp/requestid/404372541683676561917558/validate", generateRequestIdPath("404372541683676561917558", "validate"))
	assert.Equal(t, "/otp/requestid/404372541683676561917558/cancel", generateRequestIdPath("404372541683676561917558", "cancel"))
}

func TestVerifyByDestinationFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	_, err := otp.VerifyByDestination("61400000000", "12324")

	assert.Error(t, err)
}

func TestVerifyByDestinationSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.VerifyOtpResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	res, err := otp.VerifyByDestination("61400000000", "12324")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(testdata.VerifyOtpResponse(), res)
	assert.Equal(t, testdata.VerifyOtpResponse().Status, "Verified")
}

func TestVerifyByRequestIdFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	_, err := otp.VerifyByRequestId("404372541683676561917558", "12324")

	assert.Error(t, err)
}

func TestVerifyByRequestIdSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.VerifyOtpResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	res, err := otp.VerifyByRequestId("404372541683676561917558", "12324")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(testdata.VerifyOtpResponse(), res)
	assert.Equal(t, testdata.VerifyOtpResponse().Status, "Verified")
}

func TestCancelByDestinationFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	_, err := otp.CancelByDestination("61400000000")

	assert.Error(t, err)
}

func TestCancelByDestinationSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.VerifyOtpResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	res, err := otp.CancelByDestination("61400000000")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(testdata.CancelOtpResponse(), res)
	assert.Equal(t, testdata.CancelOtpResponse().Status, "Cancelled")
}

func TestCancelByRequestIdFailed(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetBadRequestResponse,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	_, err := otp.CancelByRequestId("404372541683676561917558")

	assert.Error(t, err)
}

func TestCancelByRequestIdSuccess(t *testing.T) {
	c := setup()
	mocks.ResponseJson = testdata.VerifyOtpResponseJson()
	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	otp := &Client{
		Handler: c,
		Logger:  l,
	}

	res, err := otp.CancelByRequestId("404372541683676561917558")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(testdata.CancelOtpResponse(), res)
	assert.Equal(t, testdata.CancelOtpResponse().Status, "Cancelled")
}
