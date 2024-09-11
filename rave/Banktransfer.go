package rave

import (
	// "fmt"
	"go/types"
)

//Sub Interface

type BankTransfer interface {
	BankTransfer(data BankTransferData) (error error, response map[string]interface{})
}

//Main Interface

// type BankTransfer interface {
// 	Transfer
// }

type BankTransferData struct {
	Pubkey         string      `json:"PBFPubKey"`
	Currency       string      `json:"currency"`
	Country        string      `json:"country"`
	Amount         string      `json:"amount"`
	Email          string      `json:"email"`
	Phonenumber    string      `json:"phonenumber"`
	Firstname      string      `json:"firstname"`
	Lastname       string      `json:"lastname"`
	IP             string      `json:"IP"`
	Txref          string      `json:"txRef"`
	Meta           types.Slice `json:"meta"`
	Subaccounts    types.Slice `json:"subaccounts"`
	Frequency      int         `json:"frequency"`
	IsBankTransfer bool        `json:"is_bank_transfer"`
	IsPermanent    int         `json:"is_permanent"`
	Narration      string      `json:"narration"`
	Duration       int         `json:"duration"`
	PaymentType   string      `json:"payment_type"`
}

type Banktransfers struct {
	Rave
}

func (b Banktransfers) SetupCharge(data BankTransferData) map[string]interface{} {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := b.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
		"PBFPubKey": b.GetPublicKey(),
		"client":    encryptedChargeData,
		"alg":       "3DES-24",
	}
	return queryParam
}

func (b Banktransfers) Transfer(data BankTransferData) (error error, response map[string]interface{}) {

	var url string
	url = b.GetBaseURL() + b.GetEndpoint("Banktransfer", "charge")
	if data.Txref == "" {
		data.Txref = GenerateRef()
	}
	postData := b.SetupCharge(data)

	err, response := MakePostRequest(postData, url)
	if err != nil {
		return err, noresponse
	}

	return nil, response

}
