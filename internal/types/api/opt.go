package api

type (
	SendOtp struct {
		Origin                string `json:"origin"`
		Destination           string `json:"destination"`
		Message               string `json:"message"`
		Length                int    `json:"length,omitempty"`
		CodeExpiry            int    `json:"codeExpiry,omitempty"`
		MessageExpiryDateTime string `json:"messageExpiryDateTime,omitempty"`
	}

	Otp struct {
		RequestId           string `json:"request_id"`
		Destination         string `json:"destination"`
		ValidUnitlTimestamp string `json:"validUnitlTimestamp"`
		CreatedTimestamp    string `json:"createdTimestamp"`
		LastEventTimestamp  string `json:"lastEventTimestamp"`
		Status              string `json:"status"`
	}

	verifyOtp struct {
		Code string `json:"code"`
	}
)
