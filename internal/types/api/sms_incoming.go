package api

type Campaign struct {
	ID int `json:"id"`
}

type SmsIncoming struct {
	Id          uint64    `json:"id"`
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Message     string   `json:"message"`
	DateTime    string   `json:"dateTime"`
	IsMultipart bool     `json:"isMultipart"`
	IsUnicode   bool     `json:"isUnicode,omitempty"`
	PartNumber  int      `json:"partNumber,omitempty"`
	TotalParts  int      `json:"totalParts,omitempty"`
	Campaign    Campaign `json:"campaign,omitempty"`
}

type SmsIncomingList struct {
	Total    uint32        `json:"total,omitempty"`
	Offset   uint16        `json:"offset,omitempty"`
	Limit    uint16        `json:"limit,omitempty"`
	Messages []SmsIncoming `json:"messages,omitempty"`
}
