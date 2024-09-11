package rave



type Flybuy interface {
	Bill(data FlyBuyData) (error error, response map[string]interface{})
}


type RequestsData struct {
	Country       string `json:"Country"`
	Amount        int    `json:"Amount"`
	CustomerId    string `json:"CustomerId"`
	RecurringType int    `json:"RecurringType"`
	IsAirtime     bool   `json:"IsAirtime"`
	BillerName    string `json:"BillerName"`
	Reference     string `json:"Reference"`
}
type ServicepayLoadData struct {
	Country       string         `json:"Country"`
	Amount        int            `json:"Amount"`
	CustomerId    string         `json:"CustomerId"`
	RecurringType int            `json:"RecurringType"`
	IsAirtime     bool           `json:"IsAirtime"`
	BillerName    string         `json:"BillerName"`
	Reference     string         `json:"Reference"`
	BatchRef      string         `json:"BatchReference"`
	CallBackURL   string         `json:"CallBackUrl"`
	Requests      []RequestsData `json:"Requests"`
}

type FlyBuyData struct {
	Service        string             `json:"service"`
	ServiceMethod  string             `json:"service_method"`
	ServiceVersion string             `json:"service_version"`
	ServiceChannel string             `json:"service_channel`
	ServicePayload ServicepayLoadData `json:"service_payload"`
	Seckey         string             `json:"secret_key"`
}



type Billpayment struct {
	Rave
}

func (b Billpayment) Bill(data FlyBuyData) (error error, response map[string]interface{}) {
	data.Seckey = b.GetSecretKey()
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "flybuy")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

