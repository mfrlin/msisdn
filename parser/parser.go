package parser

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
)

var calling_codes map[string]string
var mno_identifiers map[string]map[string]string

func init() {
	countries, _ := ioutil.ReadFile("./parser/resources/countries.json")
	json.Unmarshal(countries, &calling_codes)

	mnos, _ := ioutil.ReadFile("./parser/resources/mnos.json")
	json.Unmarshal(mnos, &mno_identifiers)
}

type MsisdnInfo struct {
	country_code string
	dialing_number string
	mno_identifier string
	subscriber_number string
}

func Parse_msisdn(msisdn string) (MsisdnInfo, error) {
	start := find_msisdn_start(msisdn)
	for i := start+1; i <= len(msisdn); i++ {
		dialing_number := msisdn[start:i]
		code, ok := calling_codes[dialing_number]
		if ok {
			subscriber_number := msisdn[i:]
			return MsisdnInfo{country_code: code,
				dialing_number: dialing_number,
				mno_identifier: find_mno_identifier(code, subscriber_number),
				subscriber_number: subscriber_number}, nil
		}
	}
	return MsisdnInfo{}, errors.New("Country code not found.")
}

func find_msisdn_start(msisdn string) (int) {
	if strings.HasSuffix(msisdn, "+") {
		return 1
	}
	if strings.HasSuffix(msisdn, "00") {
		return 2
	}
	return 0
}

func find_mno_identifier(country_code string, subscriber_number string) (string) {
	mnos := mno_identifiers[country_code]
	for i := 1; i <= len(subscriber_number); i++ {
		var ok bool
		mno_identifier, ok := mnos[subscriber_number[0: i]]
		if ok {
			return mno_identifier
		}
	}
	return "unknown"
}
