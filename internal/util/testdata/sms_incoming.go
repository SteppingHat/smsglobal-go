package testdata

import (
	"encoding/json"

	"github.com/smsglobal/smsglobal-go/internal/types/api"
)

// GetSmsIncomingResponse return SmsIncoming type
func GetSmsIncomingResponse() *api.SmsIncoming {
	d := GetSmsIncomingResponseJson()
	res := new(api.SmsIncoming)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func SmsIncomingListResponseJson() string {
	return `
		{
		  "total":2,
		  "offset":1,
		  "limit":20,
		  "messages":[
			{
			  "id":6088544242604429,
			  "origin":"SMSGlobal",
			  "destination":"61474000000",
			  "message":"Incoming sms",
              "dateTime":"2020-12-17 13:23:38 +0000",
			  "campaign": {
                "id": 1340232
              },
			  "isMultipart":false
			},
			{
			  "id":6298870819574735,
			  "origin":"SMSGlobal",
			  "destination":"61474000001",
			  "message":"Incoming sms 2",
			  "dateTime":"20201-12-18 10:36:29 +1000",
              "campaign": {
                "id": 1340232
              },
			  "isMultipart":false
			}
		  ]
		}
	`
}

// SmsIncomingListResponse return the sms incoming list
func SmsIncomingListResponse() *api.SmsIncomingList {
	d := SmsIncomingListResponseJson()
	res := new(api.SmsIncomingList)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func GetSmsIncomingResponseJson() string {
	return `
		{
		  "id":6088544242604429,
		  "origin":"SMSGlobal",
		  "destination":"61474000000",
		  "message":"Incoming sms",
		  "dateTime":"2020-07-30 13:23:38 +0000",
		  "campaign": {
			"id": 1340232
		  },
		  "isMultipart":false
		}
	`
}
