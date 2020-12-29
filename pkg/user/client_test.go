package user

import (
	c "github.com/smsglobal/smsglobal-go/pkg/client"
	"github.com/smsglobal/smsglobal-go/util/mocks"
	"github.com/smsglobal/smsglobal-go/util/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCreditBalanceFailed(t *testing.T) {

	client := c.New("key", "secret")

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	user := &Client{
		Handler: client,
	}
	_, err := user.CreditBalance()

	assert.Error(t, err)
}

func TestUserCreditBalanceSuccess(t *testing.T) {

	client := c.New("key", "secret")

	mocks.ResponseJson = testdata.CreditBalanceJson()

	client.HttpClient = &mocks.MockClient{
		DoFunc: mocks.GetOk,
	}

	user := &Client{
		Handler: client,
	}

	res, err := user.CreditBalance()

	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, testdata.BalanceResponse().Currency, res.Currency)
	assert.Equal(t, testdata.BalanceResponse().Balance, res.Balance)

}
