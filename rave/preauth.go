package rave

type PreauthCharge interface {
	ChargePreauth(data TokenizedChargeData) (
		error error,
		response map[string]interface{},
	)
}

type PreauthVerify interface {
	VerifyPreauth(data CardVerifyData) (
		error error,
		response map[string]interface{},
	)
}

type PreauthCapture interface {
	CapturePreauth(data CardValidateData) (
		error error,
		response map[string]interface{},
	)
}

type PreauthRefundOrVoid interface {
	RefundOrVoidPreauth(data CardVerifyData) (
		error error,
		response map[string]interface{},
	)
}

type PreauthInterface interface {
	PreauthCharge
	PreauthVerify
	PreauthCapture
	PreauthRefundOrVoid
}


type PreauthCaptureData struct {
	SecretKey      string         `json:"SECKEY"`
	Amount	       float64	      `json:"amount"`
	Flwref	       string	      `json:"flwRef"`

}

type PreauthRefundData struct {
	Flwref	       string	          `json:"ref"`
	Action	       string	          `json:"action"`
	SecretKey      string             `json:"SECKEY"`
}

type Preauth struct {
	Rave
	Card
}

func (p Preauth) ChargePreauth(data TokenizedChargeData) (error error, response map[string]interface{}) {
	data.Chargetype = "preauth"
	data.SecretKey = p.GetSecretKey()
	url := p.GetBaseURL() + p.GetEndpoint("preauth", "charge")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	
	return nil, response

}


func (p Preauth) VerifyPreauth(data CardVerifyData) (error error, response map[string]interface{}) {
	err, response := p.VerifyCard(data)

	if err != nil {
		return err, noresponse
	}

	return nil, response
	
}

func (p Preauth) CapturePreauth(data PreauthCaptureData) (error error, response map[string]interface{}) {
	data.SecretKey = p.GetSecretKey()
	url := p.GetBaseURL() + p.GetEndpoint("preauth", "capture")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response

}

func (p Preauth) RefundOrVoidPreauth(data PreauthRefundData) (error error, response map[string]interface{}) {
	data.SecretKey = p.GetSecretKey()
	url := p.GetBaseURL() + p.GetEndpoint("preauth", "refundorvoid")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	
	return nil, response

}