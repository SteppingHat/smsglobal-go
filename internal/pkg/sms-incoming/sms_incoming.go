package sms_incoming

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

var path = "/sms-incoming"

func (c *Client) Get(id string) (*api.SmsIncoming, error) {

	req, err := c.Handler.NewRequest(http.MethodGet, fmt.Sprintf(`%s/%s`, path, id), nil)
	if err != nil {
		return nil, err
	}

	sms := &api.SmsIncoming{}

	err = c.Handler.Do(req, sms)

	if err != nil {
		return nil, err
	}

	return sms, nil
}

func (c *Client) List(options map[string]string) (*api.SmsIncomingList, error) {

	log := c.Logger.Lgr.With().Str("SMS Incoming API Layer", "List").Logger()
	p := path
	// append filter options
	if len(options) != 0 {
		params := url.Values{}
		for k, v := range options {
			params.Add(k, v)
		}
		p = path + "?" + params.Encode()
	}

	log.Debug().Msgf("Path string %v", p)

	req, err := c.Handler.NewRequest(http.MethodGet, p, nil)
	if err != nil {
		return nil, err
	}

	list := &api.SmsIncomingList{}

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
