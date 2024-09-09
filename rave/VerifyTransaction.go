package rave

// import (
// 	"time"
// )

type VerifyTransaction interface {
	Verify(data VerifyTransactionData) (error error, response map[string]interface{})
}

type VerifyTransactionData struct {
	Txref  string `json:"txref"`
	Seckey string `json:"SECKEY"`
}
type Verifytransaction struct {
	Rave
}

func (v Verifytransaction) Verfiy(data VerifyTransactionData) (error error, response map[string]interface{}) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("verifytransaction", "verify")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
