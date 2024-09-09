package rave


// import (
// 	"strconv" 
// 	//"time"
// )

type Settlements struct {
	Rave
}
// type ListSettlement interface {
// 	ListSettlement(data string) (error error, response map[string]interface{})
// }

type FetchSettlement interface {
	FetchSettlement(data string) (error error, response map[string]interface{})
}

type ListSettlementData struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Page         string `json:"page"`
	Subaccountid string `json:"subaccountid"`
	Seckey       string `json:"seckey"`
}

type FetchSettlementData struct {
	Id     string `json:"id"`
	To     string `json:"to"`
	From   string `json:"from"`
	Seckey string `json:"seckey"`
}

type SettlementInterface interface {
	List
	Fetch
	
}

func (s Settlements) List(data ListSettlementData) (
	error error,
	response map[string]interface{},
) {
	queryParam := map[string]string{
		"seckey": s.GetSecretKey(),
		"from":     data.From,
		"to": data.To,
		"page": data.Page,
		"subaccountid": data.Subaccountid,
	}
	url := s.GetBaseURL() + s.GetEndpoint("settlement", "list")

	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Settlements) Fetch(data FetchSettlementData) (
	error error,
	response map[string]interface{},
) {
	queryParam := map[string]string{
		"seckey": s.GetSecretKey(),
		"id":     data.Id,
		"from": data.From,
		"to": data.To,
	}
	url := s.GetBaseURL() + s.GetEndpoint("settlement", "fetch")
	url+=data.Id
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
