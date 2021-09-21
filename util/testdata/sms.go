package testdata

import (
	"encoding/json"
	"github.com/smsglobal/smsglobal-go/types/api"
)

// GetSmsResponse return Sms type
func GetSmsResponse() *api.Sms {
	d := SentToSingleDestinationResponse()
	res := new(api.Sms)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func SmsListResponseJson() string {
	return `
		{
		  "total":4165,
		  "offset":1,
		  "limit":2,
		  "messages":[
			{
			  "id":6088544242604429,
			  "outgoing_id":5252344293,
			  "origin":"NodeSdk",
			  "destination":"61474950859",
			  "message":"Test sms from GO sdk",
			  "status":"delivered",
			  "dateTime":"2020-08-18 10:36:29 +1000"
			},
			{
			  "id":6298870819574735,
			  "outgoing_id":5252344303,
			  "origin":"NodeSdk2",
			  "destination":"61474950859",
			  "message":"Test sms from GO sdk",
			  "status":"delivered",
			  "dateTime":"2020-08-18 10:36:29 +1000"
			}
		  ]
		}
	`
}

// SmsListResponse SmsResponse return Sms type
func SmsListResponse() *api.SmsList {
	d := SmsListResponseJson()
	res := new(api.SmsList)
	_ = json.Unmarshal([]byte(d), res)

	return res
}

func SentToSingleDestinationResponse() string {
	return `
		{
		  "id":6359736682344313,
		  "outgoing_id":5211897953,
		  "origin":"NodeSdk",
		  "destination":"61474950800",
		  "message":"Test sms from GO sdk",
		  "status":"sent",
		  "dateTime":"2020-07-30 13:23:38 +0000"
		}
	`
}

func SentToMultipleDestinationsResponse() string {
	return `
		{
		  "messages":[
			{
			  "outgoing_id":5211920573,
			  "origin":"NodeSdk",
			  "destination":"61488265265",
			  "message":"Test sms from GO sdk",
			  "dateTime":"2020-07-30 14:29:50 +0000",
			  "status":"Processing"
			},
			{
			  "outgoing_id":5211920583,
			  "origin":"NodeSdk",
			  "destination":"61488265266",
			  "message":"Test sms from GO sdk",
			  "dateTime":"2020-07-30 14:29:50 +0000",
			  "status":"Processing"
			}
		  ]
		}
	`
}

// IncompleteMessageResponse when an incomplete sms sent
func IncompleteMessageResponse() string {
	return `{"status":"OK","data":{"messages":[]}}`
}
