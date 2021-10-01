package smsglobal

import (
	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	e "github.com/smsglobal/smsglobal-go/pkg/error"
	"github.com/smsglobal/smsglobal-go/pkg/logger"
	"github.com/smsglobal/smsglobal-go/internal/pkg/sms"
	"github.com/smsglobal/smsglobal-go/internal/pkg/user"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
)

var (
	lg = logger.CreateLogger(constants.DebugLevel).Lgr.With().Str("SMSGlobal", "Client").Logger()
)

// SMSGlobal defines the SMSGlobal client.
type SMSGlobal struct {
	User *user.Client
	Sms *sms.Client
}

// New Init initializes the SMSGlobal client with all available resources
func New(key, secret string) (*SMSGlobal, error) {

	lg.Info().Msgf("Creating SMSGlobal instance")

	if key == "" || secret == "" {
		return nil, &e.Error{Message: "API key and Secret are required!", Code: constants.DefaultCode}
	}

	s := new(SMSGlobal)

	s.User = &user.Client{Handler: client.New(key, secret)}

	s.Sms = &sms.Client{Handler: client.New(key, secret)}

	return s, nil
}
