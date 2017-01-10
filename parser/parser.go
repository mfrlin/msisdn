package parser

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
	"runtime"
	"path/filepath"
)

var calling_codes map[string]string
var mno_identifiers map[string](map[string]string)

func init() {
	// this is a replacement for os.path.dirname(os.path.realpath(__file__)) in python
	// TODO: it's probably a hack and should be solved in another manner
	_, current_file, _, _ := runtime.Caller(1)
	dir := filepath.Dir(current_file)
	countries, _ := ioutil.ReadFile(filepath.Join(dir, "/resources/countries.json"))
	json.Unmarshal(countries, &calling_codes)

	mnos, _ := ioutil.ReadFile(filepath.Join(dir, "/resources/mnos.json"))
	json.Unmarshal(mnos, &mno_identifiers)
}

type MsisdnInfo struct {
	country_code string
	dialing_number string
	mno_identifier string
	subscriber_number string
}

func ParseMsisdn(msisdn string) (MsisdnInfo, error) {
	start := findMsisdnStart(msisdn)
	for i := start+1; i <= len(msisdn); i++ {
		dialing_number := msisdn[start:i]
		code, ok := calling_codes[dialing_number]
		if ok {
			subscriber_number := msisdn[i:]
			return MsisdnInfo{country_code: code,
				dialing_number: dialing_number,
				mno_identifier: findMnoIdentifier(code, subscriber_number),
				subscriber_number: subscriber_number}, nil
		}
	}
	return MsisdnInfo{}, errors.New("Country code not found.")
}

func findMsisdnStart(msisdn string) int {
	if strings.HasPrefix(msisdn, "+") {
		return 1
	}
	if strings.HasPrefix(msisdn, "00") {
		return 2
	}
	return 0
}

func findMnoIdentifier(country_code string, subscriber_number string) string {
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
