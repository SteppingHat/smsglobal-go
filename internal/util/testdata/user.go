package testdata

import (
	"encoding/json"

	"github.com/smsglobal/smsglobal-go/internal/types/api"
)

func CreditBalanceJson() string {
	return `{"balance" : 15,"currency" : "EUR"}`
}

func BalanceResponse() *api.BalanceResponse {
	requestJson := CreditBalanceJson()
	res := new(api.BalanceResponse)

	_ = json.Unmarshal([]byte(requestJson), res)

	return res
}
