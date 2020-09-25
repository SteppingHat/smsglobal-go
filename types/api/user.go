package api

// BalanceResponse - Represents the response from SMSGlobal rest api for /user/credit-balance request
type BalanceResponse struct {
	Balance float64 `json:"balance"`
	Currency string `json:"currencyy; omitempty"`
}