package client

import (
	"fmt"

	"github.com/smsglobal/smsglobal-go/types/api"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"github.com/smsglobal/smsglobal-go/util/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"regexp"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	client := New("key", "secret")

	assert.Equal(t, "key", client.Key)
	assert.Equal(t, "secret", client.Secret)

	defaultClient := &http.Client{
		Timeout: constants.Timeout * time.Second,
	}

	assert.Equal(t, defaultClient, client.HttpClient)
}

func TestGenerateAuthToken(t *testing.T) {

	client := New("key", "secret")

	token := client.generateAuthToken()
	assert.NotNil(t, token)

	// assert for string format matching `MAC id="%s", ts="%d", nonce="%d", mac="%s"`
	match, _ := regexp.MatchString("MAC id=\"(.*)\", ts=\"(\\d*)\", nonce=\"(\\d*)\", mac=\"(.*)\"", token)
	assert.True(t, match)
}

func TestNewRequest(t *testing.T) {

	client := New("key", "secret")

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	req, err := client.NewRequest("POST", "/sms", `{"balance" : 5,"currency" : "AUD"}`)

	assert.NoError(t, err)
	assert.Equal(t, http.MethodPost, client.method)
	assert.Equal(t, constants.ContentType, req.Header.Get("Accept"))
	assert.Equal(t, constants.ContentType, req.Header.Get("Content-Type"))
	assert.Equal(t, "utf-8", req.Header.Get("Accept-Charset"))
	assert.NotEmpty(t, req.Header.Get("Authorization"))
}

func TestDo(t *testing.T) {

	client := New("key", "secret")

	mocks.ResponseJson = `{"balance" : 5,"currency" : "INR"}`

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	req, err := client.NewRequest("POST", "/sms", `{"balance" : 5,"currency" : "AUD"}`)

	assert.NoError(t, err)
	assert.Equal(t, client.method, http.MethodPost)
	assert.NotNil(t, req)

	balance := &api.BalanceResponse{}

	res, err := client.Do(req, balance)

	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.EqualValues(t, 5.00, balance.Balance)
	assert.EqualValues(t, "INR", balance.Currency)
}

func TestDoWithGarbageResponse(t *testing.T) {
	client := New("key", "secret")

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	req, err := client.NewRequest("POST", "/sms", `{"balance" : 5,"currency" : "AUD"}`)

	assert.NoError(t, err)
	assert.Equal(t, client.method, http.MethodPost)
	assert.NotNil(t, req)

	balance := &api.BalanceResponse{}

	res, err := client.Do(req, balance)

	//assert.IsType(t, response.HTTPResponse{}, res)

	fmt.Printf("%+v", res)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "invalid character 'g' looking for beginning of object key string", "Invalid response")

}
