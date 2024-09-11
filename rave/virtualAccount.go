package rave

// import (
// 	"time"
// )

type CreateAcct interface {
	Create(data CreateData) (error error, response map[string]interface{})
}

type CreateAcctData struct {
	Email       string `json:"email"`
	Ispermanent bool   `json:"is_permanent"`
	Frequency   int    `json:"frequency"`
	Duration    int    `json:"duration"`
	Narration   string `json:"narration"`

	Seckey string `json:"seckey"`
}
type Virtualaccount struct {
	Rave
}

func (v Virtualaccount) Create(data CreateAcctData) (error error, response map[string]interface{}) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualaccount", "virtualaccountnumber")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
