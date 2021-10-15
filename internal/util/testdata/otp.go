package testdata

import (
	"encoding/json"
	"github.com/smsglobal/smsglobal-go/internal/types/api"
)

func SendOtpResponse() *api.Otp {
	d := SendOtpResponseJson()

	res := new(api.Otp)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func SendOtpResponseJson() string {
	return `
	  {
		"requestId": "404372541683676561917558",
		"destination": "61400000000",
		"validUnitlTimestamp": "2020-11-18 17:08:14",
		"createdTimestamp": "2020-11-18 16:58:14",
		"lastEventTimestamp": "2020-11-18 16:58:14",
		"status": "Sent"
	  }
	`
}

func VerifyOtpResponse() *api.Otp {
	d := VerifyOtpResponseJson()

	res := new(api.Otp)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func VerifyOtpResponseJson() string {
	return `
	  {
		"requestId": "404372541683676561917558",
		"destination": "61400000000",
		"validUnitlTimestamp": "2020-11-18 17:08:14",
		"createdTimestamp": "2020-11-18 16:58:14",
		"lastEventTimestamp": "2020-11-18 16:58:14",
		"status": "Verified"
	  }
	`
}

func CancelOtpResponse() *api.Otp {
	d := CancelOtpResponseJson()

	res := new(api.Otp)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func CancelOtpResponseJson() string {
	return `
	  {
		"requestId": "404372541683676561917558",
		"destination": "61400000000",
		"validUnitlTimestamp": "2020-11-18 17:08:14",
		"createdTimestamp": "2020-11-18 16:58:14",
		"lastEventTimestamp": "2020-11-18 16:58:14",
		"status": "Cancelled"
	  }
	`
}
