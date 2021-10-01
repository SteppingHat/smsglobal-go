package testdata

import (
	"encoding/json"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
)

var rawJson = `{"balance" : 15,"currency" : "EUR"}`

func CreditBalanceJson() string {
	return `{"balance" : 15,"currency" : "EUR"}`
}

func BalanceResponse() *api.BalanceResponse {
	var requestJson string = CreditBalanceJson()
	res := new(api.BalanceResponse)

	_ = json.Unmarshal([]byte(requestJson), res)

	return res
}
