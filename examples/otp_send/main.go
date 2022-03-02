package otp_send

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

	otp := client.CreateOtp()
	otp.SetOrigin("From number")
	otp.SetDestination("Destination number")
	otp.SetLength(4)
	otp.SetMessage("{*code*} is your SMSGlobal verification code.")

	res, err := client.Otp.Send(otp)

	if err != nil {
		fmt.Printf("Error while sending the OTP sms: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}
