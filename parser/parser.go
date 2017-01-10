package parser

import (
	"io/ioutil"
	"encoding/json"
)

type MsisdnInfo struct {
	country_code string
}

func Parse_msisdn(msisdn string) (MsisdnInfo) {
	file, _ := ioutil.ReadFile("./parser/resources/countries.json")
	var calling_codes map[string]string
	json.Unmarshal(file, &calling_codes)
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
