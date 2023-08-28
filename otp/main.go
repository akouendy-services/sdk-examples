package main

import (
	"context"
	"fmt"
	"os"

	"github.com/akouendy-services/akouendy-sdk-go/otp"
	"github.com/google/uuid"
)

var (
	otpClient otp.Client
)

func init() {
	config := otp.Config{}
	config.WithApplication(os.Getenv("OTP_APP")).
		WithBaseUrl("https://otp-api.akouendy.com").
		//WithBaseUrl("http://localhost:9000").
		WithSecret(os.Getenv("OTP_SECRET")).
		WithDevMode(true)
	otpClient = *otp.NewClient(config)
	otpClient.DevMode()
}

func main() {
	// Get Provider list
	list, _ := otpClient.GetProviders(context.TODO())
	fmt.Println("Providers list: ", list)

	// Send otp
	userID := uuid.New()
	data := otp.InitRequest{Receiver: os.Getenv("OTP_PHONE"), ID: userID, Provider: "akouendy_sn_sms"}

	if result, err := otpClient.Init(context.TODO(), data); err != nil {
		fmt.Println("Error sending OTP")
	} else {
		// Validate code
		code := result.Data.Otp
		data := otp.ValidateRequest{ID: userID, Code: code}
		_, err := otpClient.Validate(context.Background(), data)

		if err == nil {
			fmt.Println("Code validated")
		} else {
			fmt.Println("Validation failed")
		}
	}

}
