package user

import (
	r "github.com/smsglobal/smsglobal-go/interface/response"
	"github.com/smsglobal/smsglobal-go/pkg/client"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/types/api"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"net/http"
)

var (
	lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "User Client").Logger()
)

type Client struct {
	Handler *client.Client
}

// CreditBalance method performs API request to get a user account balance
func (c *Client) CreditBalance() (*api.BalanceResponse, r.Response, error) {

	req, err := c.Handler.NewRequest(http.MethodGet, "/user/credit-balance", nil)
	if err != nil {
		return nil, nil, err
	}

	balance := &api.BalanceResponse{}
	res, err := c.Handler.Do(req, balance)

	if err != nil {
		return nil, res, err
	}

	return balance, res, nil
}
