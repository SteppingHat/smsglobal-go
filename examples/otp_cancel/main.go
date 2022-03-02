package otp_cancel

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

	res, err := client.Otp.CancelByRequestId("Request Id")

	// an otp code can be cancelled by using either request id or destination number
	//res, err := client.Otp.CancelByDestination("Destination number")

	if err != nil {
		fmt.Printf("Error while verifying the OTP: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}
