package main

import (
	"fmt"
	"os"

	"github.com/smsglobal/smsglobal-go"
)

func main() {

	client, err := smsglobal.New("YOUR API KEY", "YOUR API SECRET")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(0)
	}

	res, err := client.Otp.VerifyByRequestId("Request Id", "OTP code entered by your user")

	// an otp code can be verified by using either request id or destination number
	// res, err := client.Otp.VerifyByRequestId("Destination number", "OTP code entered by your user")

	if err != nil {
		fmt.Printf("Error while verifying the OTP: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}
