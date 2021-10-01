package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalance (t *testing.T) {
	res := BalanceResponse {
		Balance: 40.414,
		Currency: "USD",
	}

	assert.Equal(t, res.Balance, 40.414)
	assert.Equal(t, res.Currency, "USD")
}