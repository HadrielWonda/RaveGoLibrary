package rave

import (
	"strconv"
)

type Subscription struct {
	Rave
}

func (s Subscription) List() (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": s.GetSecretKey(),
	}
	url := s.GetBaseURL() + s.GetEndpoint("subscriptions", "list")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subscription) Fetch(id string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": s.GetSecretKey(),
		"id":     id,
	}
	url := s.GetBaseURL() + s.GetEndpoint("subscriptions", "fetch")
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subscription) Cancel(id int) (error error, response map[string]interface{}) {
	paymentData := struct {
		Seckey string `json:"seckey"`
	}{
		Seckey: s.GetSecretKey(),
	}

	url := s.GetBaseURL() + s.GetEndpoint("subscriptions", "cancel")
	url += strconv.Itoa(id)
	url += "/cancel"
	err, response := MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (s Subscription) Activate(id int) (error error, response map[string]interface{}) {
	paymentData := struct {
		Seckey string `json:"seckey"`
	}{
		Seckey: s.GetSecretKey(),
	}

	url := s.GetBaseURL() + s.GetEndpoint("subscriptions", "activate")
	url += strconv.Itoa(id)
	url += "/activate"
	err, response := MakePostRequest(paymentData, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
