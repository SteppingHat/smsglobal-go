package sms

import (
	"fmt"
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"net/http"
	"net/url"
)

type Client struct {
	Handler *client.Client
	Logger  *logger.Logger
}

var (
	path = "/sms"
)

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