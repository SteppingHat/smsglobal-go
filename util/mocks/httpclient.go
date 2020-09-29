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
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)

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
