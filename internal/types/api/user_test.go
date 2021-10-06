package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance (t *testing.T) {
	res := BalanceResponse {
		Balance: 40.414,
		Currency: "USD",
	}

	assert.Equal(t, res.Balance, 40.414)
	assert.Equal(t, res.Currency, "USD")
}