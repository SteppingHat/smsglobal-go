package api

import (
	"time"
)

// Message represents an outgoing sms response
type Message struct {
	Id          int64  `json:"id"`
	OutgoingId  uint64 `json:"outgoing_id,omitempty"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	DateTime    string `json:"dateTime"`
}

// SmsList struct represents the list of outgoing messages
type SmsList struct {
	Total    uint16    `json:"total,omitempty"`
	Offset   uint16    `json:"offset,omitempty"`
	Limit    uint16    `json:"limit,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

// SmsResponse represents an outgoing message response received after sending a message using either SendSingleSms or SendMultipleSms
type SmsResponse struct {
	Messages []Message `json:"messages,omitempty"`
}

// SendSingleSms represents an outgoing message for sending out
type SendSingleSms struct {
	Origin            string   `json:"origin,omitempty"`
	Destination       string   `json:"destination,omitempty"`
	Destinations      []string `json:"destinations,omitempty"`
	Message           string   `json:"message,omitempty"`
	ScheduledDateTime string   `json:"scheduledDateTime,omitempty"` // UTC time in yyyy-MM-dd HH:mm:ss
	Campaign          string   `json:"campaign,omitempty"`
	SharedPool        string   `json:"sharedPool,omitempty"`
	NotifyUrl         string   `json:"notifyUrl,omitempty"`
	IncomingUrl       string   `json:"incomingUrl,omitempty"`
	ExpiryDateTime    string   `json:"expiryDateTime,omitempty"` // UTC time in yyyy-MM-dd HH:mm:ss
}

// SendMultipleSms represents an array of Sms for sending out in one request
type SendMultipleSms struct {
	Messages []*SendSingleSms `json:"messages,omitempty"`
}

func (s *SendSingleSms) SetOrigin(d string) {
	s.Origin = d
}

func (s *SendSingleSms) SetDestination(d string) {
	s.Destination = d
}

func (s *SendSingleSms) AddDestination(d string) {
	s.Destinations = append(s.Destinations, d)
}
func (s *SendSingleSms) SetMessage(d string) {
	s.Message = d
}

func (m *SendMultipleSms) AddMessage(s *SendSingleSms) {
	m.Messages = append(m.Messages, s)
}

func (s *SendSingleSms) SetScheduledDateTime(t time.Time) {
	s.ScheduledDateTime = t.Format("2006-01-01 15:04:05")
}

func (s *SendSingleSms) SetCampaign(d string) {
	s.Campaign = d
}

func (s *SendSingleSms) SetSharedPool(d string) {
	s.SharedPool = d
}

func (s *SendSingleSms) SetNotifyUrl(d string) {
	s.NotifyUrl = d
}

func (s *SendSingleSms) SetIncomingUrl(d string) {
	s.IncomingUrl = d
}

func (s *SendSingleSms) SetExpiryDateTime(t time.Time) {
	s.ExpiryDateTime = t.Format("2006-01-01 15:04:05")
}
