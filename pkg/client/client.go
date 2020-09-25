package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	r "github.com/smsglobal/smsglobal-go/interface/response"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// TODO move to constructor
var (
	lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "HTTP Client").Logger()
)

// client defines information that can be used to make a request to SMSGlobal Rest API.
type Client struct {
	method  string
	path    string
	client  *http.Client
	baseUrl *url.URL
	timeout time.Duration
	Key     string // API KEY
	Secret  string // API SECRET
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	lg.Debug().Msgf("Creating new http request instance")

	// TODO move to constructor
	baseUrl, err := c.baseUrl.Parse(fmt.Sprintf("%s%s", constants.Host, path))

	if err != nil {
		return nil, err
	}

	// TODO move to constructor and set from their
	// set URL struct if given path is valid
	c.baseUrl = baseUrl
	c.client = &http.Client{
		Timeout: constants.Timeout * time.Second,
	}

	c.method = method
	buffer := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buffer).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.baseUrl.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", constants.ContentType)
	req.Header.Add("Accept", constants.ContentType)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Authorization", c.generateAuthToken())
	req.Header.Add("User-Agent", constants.UserAgent)

	// TODO clean up before MR
	//lg.Debug().Msgf("Authorization header: %v", req.Header.Get("Authorization"))

	return req, nil
}

// generateAuthToken Generate authorization token string for each request
func (c *Client) generateAuthToken() string {

	rand.Seed(time.Now().UnixNano())
	timestamp := int(time.Now().Unix())
	nonce := rand.Intn(1000000000)

	var protocol int

	if c.baseUrl.Scheme == "https" {
		protocol = 443
	} else {
		protocol = 80
	}

	// raw string for HMAC generation
	auth := fmt.Sprintf("%d\n%d\n%s\n%s\n%s\n%d\n\n", timestamp, nonce, c.method, c.baseUrl.Path, c.baseUrl.Host, protocol)

	// TODO clean up before MR
	lg.Debug().Msgf("Raw auth string: %v", auth)

	// generate new HMAC hash
	h := hmac.New(sha256.New, []byte(c.Secret))

	// write Data to it
	h.Write([]byte(auth))

	// Encode HMAC hash bytes to base64 string
	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf(`MAC id="%s", ts="%d", nonce="%d", mac="%s"`, c.Key, timestamp, nonce, hash)
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (r.Response, error) {

	lg.Debug().Msgf("Sending %s request to %s", c.method, c.baseUrl)

	res, err := c.client.Do(req)

	if err != nil {
		return nil, &e.Error{Message: "Failed to make a request", Code: constants.DefaultCode, Response: res}
	}

	defer res.Body.Close()

	httpResponse := r.NewResponse(res)
	err = checkResponse(res)

	if err != nil {
		return httpResponse, err
	}

	if v != nil {
		err = json.NewDecoder(res.Body).Decode(v)
	}

	lg.Debug().Msg("api request completed")

	return httpResponse, err
}

// ErrorResponse provides a message.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message;omitempty"` // error message
}

// checkResponse performs required checks whether there is any error or not
func checkResponse(r *http.Response) error {

	lg.Debug().Msgf("HTTP status code: %d", r.StatusCode)

	// a successful request status code must be between 200 and 299
	if c := r.StatusCode; http.StatusOK <= c && c < http.StatusMultipleChoices {
		return nil
	}

	errorResponse := &e.Error{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	lg.Debug().Msgf("HTTP Response body: %s", string(data))

	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
