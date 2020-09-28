package smsglobal

import (
	"github.com/smsglobal/smsglobal-go/pkg/client"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/types/constants"
	"github.com/smsglobal/smsglobal-go/user"
	"net/http"
)

var (
	lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "Client").Logger()
)

// SMSGlobal defines the SMSGlobal client.
type SMSGlobal struct {
	User *user.Client
}

// Init initializes the SMSGlobal client with all available resources
func New(key, secret string) (*SMSGlobal, error) {

	lg.Info().Msgf("Creating SMSGlobal instance")

	if key == "" || secret == "" {
		return nil, &e.Error{Message: "API key and Secret are required!"}
	}

	s := new(SMSGlobal)

	s.User = &user.Client{Handler: client.New(key, secret, http.DefaultClient)}
	return s, nil
}
