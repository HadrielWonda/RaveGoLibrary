package rave

import (
	"go/types"
)

var noresponse = map[string]interface{}{
	"": "",
}

type AccountCharge interface {
	ChargeAccount(data AccountChargeData) (error error, response map[string]interface{})
}

type AccountValidate interface {
	ValidateAccount(data AccountValidateData) (error error, response map[string]interface{})
}

type AccountVerify interface {
	VerifyAccount(data AccountVerifyData) (error error, response map[string]interface{})
}

type AccountInterface interface {
	AccountCharge
	AccountValidate
	AccountVerify
}

type AccountChargeData struct {
	Cardno                string         `json:"cardno"`
	Cvv                   string         `json:"cvv"`
	Accountbank           string         `json:"accountbank"`
	Accountnumber         string         `json:"accountnumber"`
	Paymenttype           string         `json:"payment_type"`
	Amount                float64        `json:"amount"`
	Currency              string         `json:"currency"`
	Country               string         `json:"country"`
	Bvn                   string         `json:"bvn"`
	Passcode              string         `json:"passcode"`
	CustomerPhone         string         `json:"customer_phone"`
	Firstname             string         `json:"firstname"`
	Lastname              string         `json:"lastname"`
	Email                 string         `json:"email"`
	IP                    string         `json:"IP"`
	Txref		          string	     `json:"txRef"`
	SuggestedAuth         string         `json:"suggested_auth"`
	RedirectUrl           string         `json:"redirect_url"`
	Subaccounts           types.Slice    `json:"subaccounts"`
	DeviceFingerprint     string         `json:"device_fingerprint"`
	Meta                  types.Slice    `json:"meta"`
}

type AccountValidateData struct {
	PublicKey        string           `json:"PBFPubKey"`
	Reference   string           `json:"transactionreference"`
	Otp              string           `json:"otp"`
}

type AccountVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       float64	      `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}

type Account struct {
	Rave
}

func (a Account) ChargeAccount(data AccountChargeData) (error error, response map[string]interface{}) {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := a.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": a.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
	}
	
	url := a.GetBaseURL() + a.GetEndpoint("account", "charge")

	err, response := MakePostRequest(queryParam, url)
	if err != nil {
		return err, noresponse
	}

	return nil, response

}

// Validates account charge using otp
func (a Account) ValidateAccount(data AccountValidateData) (error error, response map[string]interface{}) {
	data.PublicKey = a.GetPublicKey()
    url := a.GetBaseURL() + a.GetEndpoint("account", "validate")
    err, response := MakePostRequest(data, url)
    if err != nil {
        return err, noresponse
    }

    return nil, response
}

// Verifies the transaction, amount, and currency
func (a Account) VerifyAccount(data AccountVerifyData) (error error, response map[string]interface{}) {
	data.SecretKey = a.GetSecretKey()
    url := a.GetBaseURL() + a.GetEndpoint("account", "verify")
	err, response := MakePostRequest(data, url)
	
	transactionRef := response["data"].(map[string]interface{})["txref"].(string)
	status := response["status"].(string) 
	chargeCode := response["data"].(map[string]interface{})["chargecode"].(string)
	amount := response["data"].(map[string]interface{})["chargedamount"].(float64)
	currency := response["data"].(map[string]interface{})["currency"].(string)
	
	transactionReference := data.Reference
	currencyCode := data.Currency
	chargedAmount := data.Amount
	
	err = VerifyTransactionReference(transactionRef, transactionReference)
	err = VerifySuccessMessage(status)
	err = VerifyChargeResponse(chargeCode)
	err = VerifyCurrencyCode(currency, currencyCode)
	err = VerifyChargedAmount(amount, chargedAmount)

    if err != nil {
        return err, noresponse
    }

    return nil, response
}