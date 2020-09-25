package smsglobal

import (
	r "github.com/smsglobal/smsglobal-go/interface/response"
	c "github.com/smsglobal/smsglobal-go/pkg/client"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/types/api"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"github.com/smsglobal/smsglobal-go/types/endpoints"
	"net/http"
)

var (
	lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "API Client").Logger()
)

// SMSGlobal defines the SMSGlobal client.
type SMSGlobal struct {
	client c.Client
}

// NewClient represents the constructor for the struct SMSGlobal
func NewClient(key, secret string) (*SMSGlobal, error) {

	lg.Info().Msgf("Creating SMSGlobal instance")

	if key == "" || secret == "" {
		return nil, &e.Error{Message: "API key and Secret are required!"}
	}

	s := &SMSGlobal{
		client: c.Client{Key: key, Secret: secret},
	}

	return s, nil
}

// CreditBalance method performs API request to get a user account balance
func (s *SMSGlobal) CreditBalance() (*api.BalanceResponse, r.Response, error) {

	req, err := s.client.NewRequest(http.MethodGet, endpoints.CreditBalance, nil)
	if err != nil {
		return nil, nil, err
	}

	balance := &api.BalanceResponse{}
	res, err := s.client.Do(req, balance)

	if err != nil {
		return nil, res, err
	}

	lg.Debug().Msgf("User account balance: %v", balance)

	return balance, res, nil
}
