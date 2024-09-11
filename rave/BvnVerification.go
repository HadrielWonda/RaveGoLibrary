package rave

// import (
// 	"strconv"
// )

type BVN struct {
	Rave
}

// type Bvn interface {
// 	Withdraw(data string) (error error, response map[string]interface{}) oops!!
// }

// type BvnData struct {
// 	BvnNumber   string `json:"bvn"`
// 	Seckey string `json:"seckey"`
// }






func (b BVN) Bvn(bvn string) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
		"seckey": b.GetSecretKey(),	
	}
	url := b.GetBaseURL() + b.GetEndpoint("bvn", "bvnverification") + bvn
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
