package rave

import (
	"time"
)


///<summary>
/// I don't know how this will behave with the rampant changes to virtual card laws in Nigeria
///</summary>

type Create interface {
	Create(data CreateData) (error error, response map[string]interface{})
}

type List interface {
	List(data List) (error error, response map[string]interface{})
}

type Get interface {
	Get(data string) (error error, response map[string]interface{})
}

// type Terminate interface {
// 	Terminate(id string) (error error, response map[string]interface{})
// }

type Fund interface {
	Fund(data string) (error error, response map[string]interface{})
}

type Fetch interface {
	Fetch(data string) (error error, response map[string]interface{})
}

type Withdraw interface {
	Withdraw(data string) (error error, response map[string]interface{})
}
type Freeze interface {
	Freeze(data string) (error error, response map[string]interface{})
}

type Virtualcard interface {
	Create
	List
	Get
	Fund
	// Fetch  TBC...
	Withdraw
}

type CreateData struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	BillingName       string `json:"billing_name"`
	BillingAddress    string `json:"billing_address"`
	BillingCity       string `json:"billing_city"`
	BillingState      string `json:"billing_state"`
	BillingPostalCode string `json:"billing_postal_code"`
	BillingCountry    string `json:"billing_country"`
	Seckey            string `json:"seckey"`
	CallbackURL       string `json:"callback_url"`
}

type ListData struct {
	Page   string `json:"page"`
	Seckey string `json:"seckey"`
}

type GetData struct {
	Id     int32  `json:"page"`
	Seckey string `json:"seckey"`
}

type FundData struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	DebitCurrency string `json:"debit_currency"`
	Seckey        string `json:"seckey"`
}

type WithdrawData struct {
	CardId string `json:"card_id"`
	Amount string `json:"amount"`
	Seckey string `json:"seckey"`
}

type FreezeData struct {
	CardId       string `json:"card_id"`
	StatusAction string `json:"status_action"`
	Seckey       string `json:"seckey"`
}

type FetchData struct {
	FromDate  time.Time `json:"FromDate"`
	ToDate    time.Time `json:"ToDate"`
	PageIndex int32     `json:"PageIndex"`
	PageSize  int32     `json:"PageSize"`
	CardId    string    `json:"CardId"`
	Seckey    string    `json:"seckey"`
}

type Virtualcards struct {
	Rave
}

func (v Virtualcards) Create(data CreateData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "create")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) List(data ListData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "list")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) Get(data GetData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "get")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) Fund(data FundData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "fund")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) Withdraw(data WithdrawData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "withdraw")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) Freeze(data FreezeData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "freeze")
	url += data.CardId
	url += "/status/"
	url += data.StatusAction
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

func (v Virtualcards) Fetch(data FetchData) (
	error error,
	response map[string]interface{},
) {
	data.Seckey = v.GetSecretKey()
	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "fetch")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

