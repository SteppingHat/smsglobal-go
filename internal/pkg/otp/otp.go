package otp

import (
	"fmt"
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"net/http"
)

type (
	Client struct {
		Handler *client.Client
		Logger  *logger.Logger
	}
	verifyOtp struct {
		Code string `json:"code"`
	}
)

var path = "/otp"

// generateMsisdnPath for given destination number and request type (validate or cancel)
func generateMsisdnPath(msisdn string, requestType string) string {
	return fmt.Sprintf("%s/%s/%s", path, msisdn, requestType)
}

// generateRequestIdPath for given request id and request type (validate or cancel)
func generateRequestIdPath(id string, requestType string) string {
	return fmt.Sprintf("%s/requestid/%s/%s", path, id, requestType)
}

// Send an otp code to the recipient mobile number
func (c *Client) Send(body *api.SendOtp) (*api.Otp, error) {
	log := c.Logger.Lgr.With().Str("OTP API Layer", "Send").Logger()

	log.Debug().Msg("Sending an otp message")

	return c.do(path, body)
}

// VerifyByDestination Verify an otp code input by a user using destination number
func (c *Client) VerifyByDestination(msisdn, code string) (*api.Otp, error) {
	log := c.Logger.Lgr.With().Str("OTP API Layer", "VerifyByDestination").Logger()
	log.Debug().Msg(fmt.Sprintf("Verifying an otp code: %s using destination number: %s", code, msisdn))

	p := generateMsisdnPath(msisdn, "validate")

	return c.do(p, &verifyOtp{Code: code})
}

// VerifyByRequestId Verify an otp code input by a user using request id
func (c *Client) VerifyByRequestId(id, code string) (*api.Otp, error) {
	log := c.Logger.Lgr.With().Str("OTP API Layer", "VerifyByRequestId").Logger()
	log.Debug().Msg(fmt.Sprintf("Verifying an otp code: %s using request id: %s", code, id))

	p := generateRequestIdPath(id, "validate")

	return c.do(p, &verifyOtp{Code: code})
}

// CancelByDestination Cancel an OTP request using destination number
func (c *Client) CancelByDestination(msisdn string) (*api.Otp, error) {
	log := c.Logger.Lgr.With().Str("OTP API Layer", "CancelByDestination").Logger()

	log.Debug().Msg(fmt.Sprintf("Cancelling an otp request using destination number: %s", msisdn))

	p := generateMsisdnPath(msisdn, "cancel")

	return c.do(p, nil)
}

// CancelByRequestId Cancel an OTP request using request id
func (c *Client) CancelByRequestId(id string) (*api.Otp, error) {
	log := c.Logger.Lgr.With().Str("OTP API Layer", "CancelByRequestId").Logger()

	log.Debug().Msg(fmt.Sprintf("Cancelling an otp request using request id: %s", id))

	p := generateRequestIdPath(id, "cancel")

	return c.do(p, nil)
}

// do send request to Client and wrap response in api.Otp
func (c *Client) do(p string, v interface{}) (*api.Otp, error) {
	req, err := c.Handler.NewRequest(http.MethodPost, p, v)

	if err != nil {
		return nil, err
	}

	o := &api.Otp{}

	err = c.Handler.Do(req, o)

	if err != nil {
		return nil, err
	}

	return o, nil
}
