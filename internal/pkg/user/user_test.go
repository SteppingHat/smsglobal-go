package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smsglobal/smsglobal-go/internal/pkg/client"
	"github.com/smsglobal/smsglobal-go/internal/types/constants"
	"github.com/smsglobal/smsglobal-go/internal/util/mocks"
	"github.com/smsglobal/smsglobal-go/internal/util/testdata"
	"github.com/smsglobal/smsglobal-go/pkg/logger"

)

var l *logger.Logger

func setup()  *client.Client{

	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)

	c := client.New("key", "secret")

	c.Logger = l

	l.Debug("Setup completed")

	return c
}


func TestUserCreditBalanceFailed(t *testing.T) {

	c := setup()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	user := &Client{
		Handler: c,
		Logger: l,
	}
	_, err := user.CreditBalance()

	assert.Error(t, err)
}

func TestUserCreditBalanceSuccess(t *testing.T) {

	c := setup()

	mocks.ResponseJson = testdata.CreditBalanceJson()

	c.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	user := &Client{
		Handler: c,
		Logger: l,
	}

	res, err := user.CreditBalance()

	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.BalanceResponse().Currency, res.Currency)
	assert.Equal(t, testdata.BalanceResponse().Balance, res.Balance)

}
