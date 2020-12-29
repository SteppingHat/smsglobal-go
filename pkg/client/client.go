package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/smsglobal/smsglobal-go/interface/apiclient"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	p "path"
	"time"
)

var lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "HTTP Client").Logger()

// client defines information that can be used to make a request to SMSGlobal Rest API.
type Client struct {
	method     string
	path       string
	HttpClient apiclient.HTTPClient
	BaseURL    *url.URL
	timeout    time.Duration
	Key        string // API key
	Secret     string // API secret
}

// New returns a new api request handler
func New(key, secret string) *Client {
	baseURL, _ := url.Parse(constants.Host)

	hc := &http.Client{
		Timeout: constants.Timeout * time.Second,
	}

	c := &Client{
		HttpClient: hc,
		BaseURL:    baseURL,
		Key:        key,
		Secret:     secret,
	}

	return c
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	lg.Debug().Msgf("Creating new http request instance")

	rel, err := url.Parse(path)

	if err != nil {
		return nil, err
	}

	// append path to existing path "/v2"
	c.BaseURL.Path = p.Join(c.BaseURL.Path, rel.Path)

	u := c.BaseURL.ResolveReference(c.BaseURL)
	c.method = method

	buffer := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buffer).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", constants.ContentType)
	req.Header.Add("Accept", constants.ContentType)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Authorization", c.generateAuthToken())
	req.Header.Add("User-Agent", constants.UserAgent)

	// TODO clean up before MR
	lg.Debug().Msgf("Authorization header: %v", req.Header.Get("Authorization"))

	return req, nil
}

// generateAuthToken Generate authorization token string for each request
func (c *Client) generateAuthToken() string {

	rand.Seed(time.Now().UnixNano())
	timestamp := int(time.Now().Unix())
	nonce := rand.Intn(1000000000)

	lg.Debug().Msgf("Given PATH %+v", c.BaseURL.Path)
	// raw string for HMAC generation
	auth := fmt.Sprintf("%d\n%d\n%s\n%s\n%s\n%d\n\n", timestamp, nonce, c.method, c.BaseURL.Path, c.BaseURL.Host, 443)

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

// Do sends an API request adn the API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) error {

	lg.Debug().Msgf("Sending %s request to %s", c.method, c.BaseURL)

	res, err := c.HttpClient.Do(req)

	if err != nil {
		return &e.Error{Message: "Failed to make a request", Code: constants.DefaultCode}
	}

	defer res.Body.Close()

	err = checkResponse(res)

	if err != nil {
		return err
	}

	if v != nil {
		err = json.NewDecoder(res.Body).Decode(v)
	}

	lg.Debug().Msg("HTTP request done")

	return err
}

// checkResponse performs required checks whether there is any error or not
func checkResponse(r *http.Response) error {
	lg.Debug().Msgf("HTTP status code: %d", r.StatusCode)

	// a successful request status code must be between 200 and 299
	if c := r.StatusCode; http.StatusOK <= c && c < http.StatusMultipleChoices {
		return nil
	}

	errorResponse := &e.Error{
		Code: r.StatusCode,
	}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil && data != nil {
		err = json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}
	return errorResponse
}
