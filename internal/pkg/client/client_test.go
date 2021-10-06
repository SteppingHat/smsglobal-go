package client

import (
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"github.com/smsglobal/smsglobal-go/internal/util/mocks"
	"github.com/smsglobal/smsglobal-go/internal/util/testdata"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
)

var l *logger.Logger

func setup() *Client{

	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)
	l.Debug("Setup completed")

	c := New("key", "secret")

	c.Logger = l

	return c
}
func TestNew(t *testing.T) {

	c := setup()
	assert.Equal(t, "key", c.Key)
	assert.Equal(t, "secret", c.Secret)

	defaultClient := &http.Client{
		Timeout: constants.Timeout * time.Second,
	}

	assert.Equal(t, defaultClient, c.HttpClient)
}

func TestGenerateAuthToken(t *testing.T) {

	c := setup()
	token := c.generateAuthToken()
	assert.NotNil(t, token)

	// assert for string format matching `MAC id="%s", ts="%d", nonce="%d", mac="%s"`
	match, _ := regexp.MatchString("MAC id=\"(.*)\", ts=\"(\\d*)\", nonce=\"(\\d*)\", mac=\"(.*)\"", token)
	assert.True(t, match)
}

func TestNewRequest(t *testing.T) {

	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	req, err := c.NewRequest(http.MethodGet, "/sms", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, c.method)
	assert.Equal(t, constants.ContentType, req.Header.Get("Accept"))
	assert.Equal(t, constants.ContentType, req.Header.Get("Content-Type"))
	assert.Equal(t, "utf-8", req.Header.Get("Accept-Charset"))
	assert.NotEmpty(t, req.Header.Get("Authorization"))
}

func TestDo(t *testing.T) {

	c := setup()

	mocks.ResponseJson = testdata.SentToSingleDestinationResponse()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	p := `{ "origin":"NodeSdk", "destination":"61474950800", "message":"Test sms from GO sdk"}`
	req, err := c.NewRequest(http.MethodPost, "/sms", p)

	assert.NoError(t, err)
	assert.Equal(t, c.method, http.MethodPost)
	assert.NotNil(t, req)

	sms := &api.Sms{}
	err = c.Do(req, sms)

	assert.NoError(t, err)
	assert.EqualValues(t, testdata.GetSmsResponse().Origin, sms.Origin)
	assert.EqualValues(t, testdata.GetSmsResponse().Destination, sms.Destination)
	assert.EqualValues(t, testdata.GetSmsResponse().Message, sms.Message)
}

func TestDoWithGarbageResponse(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	p := `{ "origin":"NodeSdk", "destination":"61474950800", "message":"Test sms from GO sdk"}`
	req, err := c.NewRequest(http.MethodPost, "/sms", p)

	assert.NoError(t, err)
	assert.Equal(t, c.method, http.MethodPost)
	assert.NotNil(t, req)

	balance := &api.BalanceResponse{}

	err = c.Do(req, balance)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character 'g' looking for beginning of object key string", "Invalid response")
}

func TestAuthenticationFailure(t *testing.T) {
	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetUnknownAuthenticationError,
	}

	p := `{ "origin":"NodeSdk", "destination":"61474950800", "message":"Test sms from GO sdk"}`
	req, err := c.NewRequest(http.MethodPost, "/sms", p)

	err = c.Do(req, new(api.BalanceResponse))

	assert.Error(t, err)
	assert.Equal(t, err.Error(), `{"code":403,"message":"Unknown Authentication Error"}`, "Invalid response")
}

func TestDoNoContentResponse(t *testing.T) {

	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetNoContent,
	}

	req, err := c.NewRequest(http.MethodDelete, "/sms/6746514019161950", nil)

	assert.NoError(t, err)
	assert.Equal(t, c.method, http.MethodDelete)
	assert.NotNil(t, req)

	err = c.Do(req, nil)
	assert.NoError(t, err)
}
