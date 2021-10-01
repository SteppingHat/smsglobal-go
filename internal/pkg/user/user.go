package user

import (
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"net/http"
)

var (
	lg   = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "User Client").Logger()
	path = "/user/credit-balance"
)

type Client struct {
	Handler *client.Client
}

// CreditBalance method performs API request to get a user account balance
func (c *Client) CreditBalance() (*api.BalanceResponse, error) {

	req, err := c.Handler.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	balance := &api.BalanceResponse{}
	err = c.Handler.Do(req, balance)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
