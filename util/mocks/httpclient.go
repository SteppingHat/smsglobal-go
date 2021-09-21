package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MockClient is the mock client
type MockClient struct {
	Json   string
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	ResponseJson string
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func GetOk(*http.Request) (*http.Response, error) {

	if ResponseJson == "" {
		ResponseJson = `{"success": true}`
	}

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(ResponseJson)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

func GetGarbageResponse(*http.Request) (*http.Response, error) {
	json := `{garbage`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

func GetUnknownAuthenticationError(*http.Request) (*http.Response, error) {

	json := `Unknown Authentication Error`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusForbidden,
		Body:       r,
	}, nil
}

// GetNoContent 204 response
func GetNoContent(*http.Request) (*http.Response, error) {
	var bodyBytes []byte

	//  http.Response guarantees that Response.Body will not be nil even if there is no response data so do the mock response
	r := ioutil.NopCloser(bytes.NewBuffer([]byte(bodyBytes)))

	return &http.Response{
		StatusCode: http.StatusNoContent,
		Body: r,
	}, nil
}

func GetNotFound(*http.Request) (*http.Response, error) {

	json := `{ "message": "" }`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       r,
	}, nil
}