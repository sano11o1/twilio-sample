package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	// from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	serviceSid := os.Getenv("VERIFY_SERVICE_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	var inputCode string

	fmt.Print("Enter Code > ")

	inputCodeError, err := fmt.Scanln(&inputCode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inputCodeError)
	}

	checkParams := &openapi.CreateVerificationCheckParams{}
	checkParams.SetTo(to)
	checkParams.SetCode(inputCode)

	checkResp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, checkParams)
	if err != nil {
		fmt.Println(err)
	} else if *checkResp.Status == "approved" {
		fmt.Println("認証成功")
	} else {
		fmt.Println("認証失敗")
	}
}
