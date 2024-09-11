package rave

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"runtime"
// 	"math/rand"
// )

type otp interface {
	otp(data OTPData) (
		error error,
		response map[string]interface{},
	)
}

type CustomerInfoData struct {
	Firstname string `json:"first_name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}

// type MediumData struct {
// 	Whatsapp string `json:"whatsapp"`
// 	SMS      string `json:"sms"`     TBC................................................................
// 	Email    string `json:"email"`
// }
type ServicepayData struct {
	LenghtOfOTP       int              `json:"length_of_otp"`
	SendOTPasCustomer bool             `json:"send_otp_to_customer"`
	CustomerInfo      CustomerInfoData `json:"customer_info"`
	Medium            []string         `json:"medium"`
	TransactionRef  string             `json:"transaction_reference"`
}

type OTPData struct {
	Service            string         `json:"service"`
	ServiceMethod      string         `json:"service_method"`
	ServiceVersion     string         `json:"service_version"`
	ServiceChannel     string         `json:"service_channel`
	ServicePayload     ServicepayData `json:"service_payload"`
	Seckey             string         `json:"secret_key"`
	OtpExpiresInMins   int            `json:"otp_expires_in_minutes"`
	SenderBusinessname string         `json:"sender_business_name"`
	SenderSameOTP      bool           `json:"send_same_otp"`
	SenderAsIs         int            `json:"send_as_is"`
}

type FlutterwaveOTP struct {
	Rave
}

func (o FlutterwaveOTP) Otp(data OTPData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = o.GetSecretKey()
	url := o.GetBaseURL() + o.GetEndpoint("flutterwaveOTP", "otp")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
