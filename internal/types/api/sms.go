package api

// Message represents an outgoing sms response
type Message struct {
	Id			int64 `json:"id"`
	OutgoingId	uint64 `json:"outgoing_id,omitempty"`
	Origin		string `json:"origin"`
	Destination string `json:"destination"`
	Message		string `json:"message"`
	Status		string 	`json:"status"`
	DateTime	string `json:"dateTime"`
}

// SmsList struct represents the list of outgoing messages
type SmsList struct {
	Total 		uint16 `json:"total,omitempty"`
	Offset 		uint16 `json:"offset,omitempty"`
	Limit 		uint16 `json:"limit,omitempty"`
	Messages 	[]Message `json:"messages,omitempty"`
}

// SmsResponse represents an outgoing message response received after sending a message using either SendSingleSms or SendMultipleSms
type SmsResponse struct {
	Messages []Message `json:"messages,omitempty"`
}

// SendSingleSms represents an outgoing message for sending out
type SendSingleSms struct {
	Origin string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
	Destinations []string `json:"destinations,omitempty"`
	Message string `json:"message,omitempty"`
}

// SendMultipleSms represents an array of Sms for sending out in one request
type SendMultipleSms struct {
	Messages 		[]*SendSingleSms `json:"messages,omitempty"`
}

func (s *SendSingleSms) SetOrigin (o string) {
	s.Origin = o
}

func (s *SendSingleSms) SetDestination(d string)  {
	s.Destination = d
}

func (s *SendSingleSms) AddDestination(d string) {
	s.Destinations = append(s.Destinations, d)
}
func (s *SendSingleSms) SetMessage(m string)  {
	s.Message = m
}

func (s *SendMultipleSms) AddMessage(sms *SendSingleSms)  {
	s.Messages = append(s.Messages, sms)
}
