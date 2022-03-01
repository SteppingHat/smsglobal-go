package main

import (
	"fmt"
	"os"

	"github.com/smsglobal/smsglobal-go"
)

func main() {
	SendOne()
	SmsList()
	SmsIncomingList()
}

func Init() *smsglobal.SMSGlobal {
	c, err := smsglobal.New("YOUR API KEY", "YOUR API SECRET")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(0)
	}
	return c
}

func SendOne() {
	client := Init()

	sms := client.CreateSms()
	sms.SetOrigin("From number")

	sms.AddDestination("Destination number 1")
	sms.AddDestination("Destination number 2")
	sms.AddDestination("Destination number 3")

	sms.SetMessage("This is a test message")

	res, err := client.Sms.SendOne(sms)

	if err != nil {
		fmt.Printf("Error while sending the message: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res.Messages)
	fmt.Println()
}

func SendMultiple() {

	client := Init()

	sms := client.CreateMultipleSms()

	// first message
	s := client.CreateSms()
	s.SetOrigin("From Number")
	s.SetMessage("First test message.")
	s.SetDestination("Destination number")

	sms.AddMessage(s)

	// second message
	s = client.CreateSms()
	s.SetOrigin("From Number")
	s.SetMessage("Second test message.")
	s.SetDestination("Destination number")

	sms.AddMessage(s)

	// second message
	s = client.CreateSms()
	s.SetOrigin("From Number")
	s.SetMessage("Third test message.")
	s.SetDestination("Destination number")

	sms.AddMessage(s)

	res, err := client.Sms.SendMultiple(sms)

	if err != nil {
		fmt.Printf("Error while sending multiple messages: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res.Messages)
	fmt.Println()
}

func SmsList() {
	options := make(map[string]string)
	options["offset"] = "1"
	options["limit"] = "5"
	client := Init()
	res, err := client.Sms.List(options)

	if err != nil {
		fmt.Printf("Error while fetching the list of mesages: %s", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Total sms found %d", res.Total)
	fmt.Println()
	fmt.Printf("Reponse received: %+v", res.Messages)

	// Loop over structs an
	for i := range res.Messages {

		fmt.Printf("#%d Outgoing Id = %v, Message = %v", i, res.Messages[i].OutgoingId, res.Messages[i].Message)
		fmt.Println()
	}
}

func SmsGet() {
	client := Init()

	res, err := client.Sms.Get("Outgoing Id or Id")

	if err != nil {
		fmt.Printf("Error while fetching the message details: %s", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()

}

func SmsDelete() {
	client := Init()

	err := client.Sms.Delete("Outgoing Id or Id")

	if err != nil {
		fmt.Printf("Error while deleting the message: %s", err.Error())
		fmt.Println()
		os.Exit(0)
	}
}

func SmsIncomingList() {
	options := make(map[string]string)
	options["offset"] = "1"
	options["limit"] = "5"
	client := Init()
	res, err := client.SmsIncoming.List(options)

	if err != nil {
		fmt.Printf("Error while fetching the list of incoming mesages: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Total sms found %d", res.Total)
	fmt.Println()
	fmt.Printf("Reponse received: %+v", res.Messages)

	// Loop over structs an
	for i := range res.Messages {

		fmt.Printf("#%d Id = %v, Message = %v", i, res.Messages[i].Id, res.Messages[i].Message)
		fmt.Println()
	}
}

func SmsIncomingGet() {
	client := Init()

	res, err := client.SmsIncoming.Get("Id")

	if err != nil {
		fmt.Printf("Error while fetching the incoming message details: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}

func SmsIncomingDelete() {
	client := Init()

	err := client.SmsIncoming.Delete("Id")

	if err != nil {
		fmt.Printf("Error while deleting the incoming message: %s \n", err.Error())
		fmt.Println()
		os.Exit(0)
	}
}
