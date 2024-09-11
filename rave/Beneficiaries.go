package rave

// import (
// 	"strconv"
// )

type TransRecipients struct {
	Rave
}

type FetchRecipientsData struct {
	SecKey              string `json:"seckey"`
	Id                  string `json:"id"`
	FullnameOrAccountNo string `json:"q"`
}

type CreateRecipientData struct {
	AccountNo string `json:"account_number"`
	Acctbank  string `json:"account_bank"`
	Seckey    string `json:"seckey"`
}

type DeleteRecipientData struct {
	Id     string `json:"id"`
	Seckey string `json:"seckey"`
}

func (t TransRecipients) List() (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": t.GetSecretKey(),
	}
	url := t.GetBaseURL() + t.GetEndpoint("Beneficiaries", "list")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (t TransRecipients) Fetch(data FetchRecipientsData) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": t.GetSecretKey(),
		"id":     data.Id,
		"q":      data.FullnameOrAccountNo,
	}
	url := t.GetBaseURL() + t.GetEndpoint("Beneficiaries", "fetch")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (t TransRecipients) CreateRecipient(data CreateRecipientData) (error error, response map[string]interface{}) {
	data.Seckey = t.GetSecretKey()
	url := t.GetBaseURL() + t.GetEndpoint("Beneficiaries", "create")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (t TransRecipients) DeleteRecipient(data DeleteRecipientData) (error error, response map[string]interface{}) {
	data.Seckey = t.GetSecretKey()
	url := t.GetBaseURL() + t.GetEndpoint("Beneficiaries", "delete")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

