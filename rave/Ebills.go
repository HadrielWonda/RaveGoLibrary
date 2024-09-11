package rave

// import (
// 	"time"
// )

type CreateOrder interface {
	Create(data CreateOrderData) (error error, response map[string]interface{})
}

type UpdateOrder interface {
	Update(data UpdateOrderData) (error error, response map[string]interface{})
}

type CreateOrderData struct {
	Currency      string `json:"currency"`
	NumberOfUnits int    `json:"numberofunits"`
	Amount        int    `json:"amount"`
	Narration     string `json:"narration"`
	Seckey        string `json:"SECKEY"`
	PhoneNumber   string `json:"phonenumber"`
	Email         string `json:"email"`
	Txref         string `json:"txRef"`
	IP            string `json:"IP"`
	Country       string `json:"country"`
}

type UpdateOrderData struct {
	Currency string `json:"currency"`
	FlwRef   string `json:"reference"`
	Amount   int    `json:"amount"`
	Seckey   string `json:"SECKEY"`
}

type Ebills struct {
	Rave
}

func (e Ebills) CreateOrder(data CreateOrderData) (error error, response map[string]interface{}) {
	data.Seckey = e.GetSecretKey()
	url := e.GetBaseURL() + e.GetEndpoint("ebills", "createorder")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (e Ebills) UpdateOrder(data UpdateOrderData) (error error, response map[string]interface{}) {
	data.Seckey = e.GetSecretKey()
	url := e.GetBaseURL() + e.GetEndpoint("ebills", "updateorder")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
