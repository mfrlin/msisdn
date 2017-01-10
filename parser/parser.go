package parser

import (
	"io/ioutil"
	"encoding/json"
)

var calling_codes map[string]string

func init() {
	file, _ := ioutil.ReadFile("./parser/resources/countries.json")
	json.Unmarshal(file, &calling_codes)
}

type MsisdnInfo struct {
	country_code string
}

func Parse_msisdn(msisdn string) (MsisdnInfo) {
	var code string
	for i := 1; i <= len(msisdn); i++ {
		var ok bool
		code, ok = calling_codes[msisdn[0:i]]
		if ok {
			break
		}
	}
	return MsisdnInfo{country_code: code}
}
