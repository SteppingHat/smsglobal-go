package api

import (
	"time"

	"github.com/smsglobal/smsglobal-go/internal/types/constants"
)

type (
	SendOtp struct {
		Origin                string `json:"origin,omitempty"`
		Destination           string `json:"destination"`
		Message               string `json:"message"`
		Length                int    `json:"length,omitempty"`
		CodeExpiry            int    `json:"codeExpiry,omitempty"`
		MessageExpiryDateTime string `json:"messageExpiryDateTime,omitempty"`
	}

	Otp struct {
		RequestId           string `json:"requestId"`
		Destination         string `json:"destination"`
		ValidUntilTimestamp string `json:"validUnitlTimestamp"`
		CreatedTimestamp    string `json:"createdTimestamp"`
		LastEventTimestamp  string `json:"lastEventTimestamp"`
		Status              string `json:"status"`
	}

	verifyOtp struct {
		Code string `json:"code"`
	}
)

func (o *SendOtp) SetOrigin(or string) {
	o.Origin = or
}

func (o *SendOtp) SetDestination(d string) {
	o.Destination = d
}

func (o *SendOtp) SetMessage(m string) {
	o.Message = m
}

func (o *SendOtp) SetLength(l int) {
	o.Length = l
}

func (o *SendOtp) SetCodeExpiry(e int) {
	o.CodeExpiry = e
}

func (o *SendOtp) SetMessageExpiryDateTime(t time.Time) {
	o.MessageExpiryDateTime = t.Format(constants.DateTimeFormat)
}
