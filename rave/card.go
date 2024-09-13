package rave

import (
	"go/types"
)


type CardCharge interface {
	ChargeCard(data CardChargeData) (error error, response map[string]interface{})
}

type CardValidate interface {
	ValidateCard(data CardValidateData) (error error, response map[string]interface{})
}

type CardVerify interface {
	VerifyCard(data CardVerifyData) (error error, response map[string]interface{})
}

type CardTokenized interface {
	TokenizedCharge(data TokenizedChargeData) (error error, response map[string]interface{})
}

type CardInterface interface {
	CardCharge
	CardValidate
	CardVerify
	CardTokenized
}

type CardChargeData struct {
	Cardno               string         `json:"cardno"`
	Cvv                  string         `json:"cvv"`
	Expirymonth          string         `json:"expirymonth"`
	Expiryyear           string         `json:"expiryyear"`
	Pin                  string         `json:"pin"`
	Amount               float64        `json:"amount"`
	Currency             string         `json:"currency"`
	Country              string         `json:"country"`
	CustomerPhone        string         `json:"customer_phone"`
	Firstname            string         `json:"firstname"`
	Lastname             string         `json:"lastname"`
	Email                string         `json:"email"`
	Ip                   string         `json:"IP"`
	Txref		         string	        `json:"txRef"`
	RedirectUrl          string         `json:"redirect_url"`
	Subaccounts          types.Slice    `json:"subaccounts"`
	DeviceFingerprint    string         `json:"device_fingerprint"`
	Meta                 types.Slice    `json:"meta"`
	SuggestedAuth        string         `json:"suggested_auth"`
	BillingZip           string         `json:"billingzip"`
	BillingCity          string         `json:"billingcity"`
	BillingAddress       string         `json:"billingaddress"`
	BillingState         string         `json:"billingstate"`
	BillingCountry       string         `json:"billingcountry"`
	Chargetype		     string	        `json:"charge_type"`

}

type CardValidateData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"`
}

type CardVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       float64	      `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}

type TokenizedChargeData struct {
	SecretKey      string         `json:"SECKEY"`
	Currency       string         `json:"currency"`
	Token          string         `json:"token"`
	Country        string         `json:"country"`
	Amount	       float64	      `json:"amount"`
	Email          string         `json:"email"`
	Firstname      string         `json:"firstname"`
	Lastname       string         `json:"lastname"`
	Ip             string         `json:"IP"`
	Txref		   string	      `json:"txRef"`
	Chargetype	  string	        `json:"charge_type"`
}

type Card struct {
	Rave
}

func (c Card) SetupCharge(data CardChargeData) map[string]interface{} {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := c.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": c.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	return queryParam
}

func (c Card) ChargeCard(data CardChargeData) (error error, response map[string]interface{}) {
	var url string
	if (data.Txref == "") {
		data.Txref = GenerateRef()
	}
	postData := c.SetupCharge(data)

	if data.Chargetype == "preauth" {
		url = c.GetBaseURL() + c.GetEndpoint("preauth", "charge")
	} else {
		url = c.GetBaseURL() + c.GetEndpoint("card", "charge")
	}
	
	err, response := MakePostRequest(postData, url)
	if err != nil {
		return err, noresponse
	}

	suggestedAuth := response["data"].(map[string]interface{})["suggested_auth"]
	if (suggestedAuth == "PIN") {
		data.SuggestedAuth = "PIN"
		postData = c.SetupCharge(data)
		err, response = MakePostRequest(postData, url)
		if err != nil {
			return err, noresponse
		}
	} else if (suggestedAuth == "AVS_VBVSECURECODE") {
		data.SuggestedAuth = "AVS_VBVSECURECODE"
		postData = c.SetupCharge(data)
		err, response = MakePostRequest(postData, url)
		if err != nil {
			return err, noresponse
		}
	}

	return nil, response

}

func (c Card) ValidateCard(data CardValidateData) (error error, response map[string]interface{}) {
	data.PublicKey = c.GetPublicKey()
	url := c.GetBaseURL() + c.GetEndpoint("card", "validate")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response

}

func (c Card) VerifyCard(data CardVerifyData) (error error, response map[string]interface{}) {
	data.SecretKey = c.GetSecretKey()
	url := c.GetBaseURL() + c.GetEndpoint("card", "verify")
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

func (c Card) TokenizedCharge(data TokenizedChargeData) (error error, response map[string]interface{}) {
	data.SecretKey = c.GetSecretKey()
	url := c.GetBaseURL() + c.GetEndpoint("card", "chargeSavedCard")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	
	return nil, response

}

