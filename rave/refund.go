// might be unstable

package rave

type RefundCharge interface {
	RefundTransaction(data RefundData) (error error, response map[string]interface{})
}

type RefundInterface interface {
	RefundCharge
}

type RefundData struct {
	Ref		       string	      `json:"ref"`
	Amount         int            `json:"amount"`
	SecretKey      string         `json:"seckey"`
}

type Refund struct {
	Rave
}

func (r Refund) RefundTransaction(data RefundData) (error error, response map[string]interface{}) {
	data.SecretKey = r.GetSecretKey()
	url := r.GetBaseURL() + r.GetEndpoint("refund", "refund")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response

}
