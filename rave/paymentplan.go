package rave

import (
	"strconv"
)

type PaymentPlanData struct {
	Amount   string `json:"amount"`
	Name     string `json:"name"`
	Interval string `json:"interval"`
	Duration string `json:"duration"`
	Seckey   string `json:"seckey"`
}

type PaymentPlan struct {
	Rave
}

func (p PaymentPlan) Create(data PaymentPlanData) (error error, response map[string]interface{}) {
	data.Seckey = p.GetSecretKey()
	url := p.GetBaseURL() + p.GetEndpoint("payment_plan", "create")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (p PaymentPlan) List() (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": p.GetSecretKey(),
	}
	url := p.GetBaseURL() + p.GetEndpoint("payment_plan", "list")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (p PaymentPlan) Fetch(id string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": p.GetSecretKey(),
		"id":     id,
	}
	url := p.GetBaseURL() + p.GetEndpoint("payment_plan", "fetch")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (p PaymentPlan) Cancel(id int) (error error, response map[string]interface{}) {
	paymentData := struct {
		Seckey string `json:"seckey"`
	}{
		Seckey: p.GetSecretKey(),
	}

	url := p.GetBaseURL() + p.GetEndpoint("payment_plan", "cancel")
	url += strconv.Itoa(id)
	url += "/cancel"
	err, response := MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (p PaymentPlan) Edit(id int, name string, status string) (error error, response map[string]interface{}) {
	paymentData := struct {
		name   string `json:"name"`
		status string `json:"status"`
		Seckey string `json:"seckey"`
	}{
		name:   name,
		status: status,
		Seckey: p.GetSecretKey(),
	}

	url := p.GetBaseURL() + p.GetEndpoint("payment_plan", "edit")
	url += strconv.Itoa(id)
	url += "/edit"
	err, response := MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
