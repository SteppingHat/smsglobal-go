# SMSGlobal Go SDK

---
![Go](https://github.com/smsglobal/smsglobal-go/workflows/Go/badge.svg?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/smsglobal/smsglobal-go/-/badge.svg)](https://sourcegraph.com/github.com/smsglobal/smsglobal-go?badge)
![Go](https://github.com/smsglobal/smsglobal-go/workflows/Go/badge.svg?branch=master&event=status)
---

The official [SMSGlobal](https://www.smsglobal.com?utm_source=dev&utm_medium=github&utm_campaign=go_sdk) Go client library.

Sign up for a [free SMSGlobal account](https://www.smsglobal.com/mxt-sign-up/?utm_source=dev&utm_medium=github&utm_campaign=go_sdk) today and get your API Key from our advanced SMS platform, MXT. Plus, enjoy unlimited free developer sandbox testing to try out your API in full!



## Install Prerequisites
The following are the prerequisites for this package

### Go v1.14 or higher

Use following page to install go https://golang.org/doc/install?download=go1.14.2.linux-amd64.tar.gz

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/smsglobal/smsglobal-go
```
---

## Usage

Import the library into your package

```go
import (
"github.com/smsglobal/smsglobal-go"
)
```

#### Send SMS

```go
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

```

#### Send OTP

```go
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

```

#### Verify OTP

The OTP code entered by your user can be verified by either using `requestId` or `destination number`. The followings are examples of each method:

```go
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

  if err != nil {
		fmt.Printf("Error while verifying the OTP: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}

```

Use `VerifyByDestination` method in order to verify the OTP code using destination number.

#### Cancel OTP

The OTP request can be cancelled if an OTP is not expired and verified yet. It can be done by either using `requestId` or `destination number`. The followings are examples of each method:

```go
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

	res, err := client.Otp.CancelByRequestId("Request Id")

	if err != nil {
		fmt.Printf("Error while verifying the OTP: %s \n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Reponse received: %+v", res)
	fmt.Println()
}

```

Use `CancelByDestination` method in order to cancel the OTP request using destination number

---

### Examples

Checkout [examples](examples) folder. It contains an example of each method available using this client library

---

## Available REST API Resources
* User
* Sms
* OTP (beta)

---

## Developers

### Installing GoDoc for Go v1.1.4
Once Go v1.14 has been installed, execute the following commands to install GoDoc from smsglobal-go project's root directory

```sh
GOBIN=`pwd`/bin && mkdir bin && go get -u golang.org/x/tools/cmd/godoc
```

Make sure you use the godoc tool that is installed in smsglobal-go/bin directory

### Running documentation server

```sh
./dev-bin/docs
```

Navigate here to see the documentation in your local browser: 
http://localhost:6060/pkg/github.com/smsglobal/smsglobal-go

### Running Tests
To run unit tests execute the following command

```sh
./dev-bin/test
```
