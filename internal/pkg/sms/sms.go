package sms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
)

type Client struct {
	Handler *client.Client
	Logger  *logger.Logger
}

var path = "/sms"

func (c *Client) SendOne(params *api.SendSingleSms) (*api.SmsResponse, error) {
	log := c.Logger.Lgr.With().Str("SMS API Layer", "SendOne").Logger()

	log.Debug().Msg("Sending a single message")

	return c.send(params)
}

func (c *Client) SendMultiple(params *api.SendMultipleSms) (*api.SmsResponse, error) {
	log := c.Logger.Lgr.With().Str("SMS API Layer", "SendMultiple").Logger()
	log.Debug().Msg("Sending multiple messages")

	return c.send(params)
}

func (c *Client) send(body interface{}) (*api.SmsResponse, error) {
	req, err := c.Handler.NewRequest(http.MethodPost, path, body)

	if err != nil {
		return nil, err
	}

	sms := &api.SmsResponse{}

	err = c.Handler.Do(req, sms)

	if err != nil {
		return nil, err
	}

	return sms, nil
}

func (c *Client) Get(id string) (*api.Sms, error) {

	req, err := c.Handler.NewRequest(http.MethodGet, fmt.Sprintf(`%s/%s`, path, id), nil)
	if err != nil {
		return nil, err
	}

	sms := &api.Sms{}

	err = c.Handler.Do(req, sms)

	if err != nil {
		return nil, err
	}

	return sms, nil
}

func (c *Client) List(options map[string]string) (*api.SmsList, error) {

	log := c.Logger.Lgr.With().Str("SMS API Layer", "List").Logger()

	// append filter options
	if len(options) != 0 {
		params := url.Values{}
		for k, v := range options {
			params.Add(k, v)
		}
		path = path + "?" + params.Encode()
	}

	log.Debug().Msgf("Path string %v", path)

	req, err := c.Handler.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	list := &api.SmsList{}

	err = c.Handler.Do(req, list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Client) Delete(id string) error {

	req, err := c.Handler.NewRequest(http.MethodDelete, fmt.Sprintf(`%s/%s`, path, id), nil)
	if err != nil {
		return err
	}

	err = c.Handler.Do(req, nil)

	if err != nil {
		return err
	}

	return nil
}
