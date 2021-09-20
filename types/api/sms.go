package api


type message struct {
	Id int64 `json:"id"`
	OutgoingId uint64 `json:"outgoing_id; omitempty"`
	Origin string `json:"origin"`
	Destination string `json:"destination"`
	Message string `json:"message"`
	Status string 	`json:"status"`
	DateTime string `json:"dateTime"`
}

// SmsList struct represents the list of outgoing messages
type SmsList struct {
	Total uint16 `json:"total; omitempty"`
	Offset uint16 `json:"offset; omitempty"`
	Limit uint16 `json:"limit; omitempty"`
	Messages []message
}


// Sms represents an outgoing messages
type Sms struct {
	message
}