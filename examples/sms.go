package main

import (
	"fmt"
	"os"

	"github.com/smsglobal/smsglobal-go"
)

func main() {
	SendOne()
	SendMultiple()
	Get()
	List()
	Delete()
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
		fmt.Printf("Error while sending the sms: %s \n", err.Error())
		return
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
		fmt.Printf("Error while sending the sms: %s \n", err.Error())
		os.Exit(0)
		return
	}

	fmt.Printf("Reponse received: %+v", res.Messages)
	fmt.Println()
}

func List() {
	options := make(map[string]string)
	options["offset"] = "1"
	options["limit"] = "5"
	client := Init()
	res, err := client.Sms.List(options)

	if err != nil {
		fmt.Printf("Error while sending the sms: %s", err.Error())
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

func Get() {
	client := Init()

	res, err := client.Sms.Get("outgoing id or id")

	if err != nil {
		fmt.Printf("Error while sending the sms: %s", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()

}

func Delete() {
	client := Init()

	err := client.Sms.Delete("outgoing id or id")

	if err != nil {
		fmt.Printf("Error while sending the sms: %s", err.Error())
		fmt.Println()
		os.Exit(0)
	}
}
