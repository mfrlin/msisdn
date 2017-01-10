package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

var callingCodes map[string]string
var mnoIdentifiers map[string](map[string]string)

func init() {
	// this is a replacement for os.path.dirname(os.path.realpath(__file__)) in python
	// TODO: it's probably a hack and should be solved in another manner
	_, currentFile, _, _ := runtime.Caller(1)
	dir := filepath.Dir(currentFile)
	countries, _ := ioutil.ReadFile(filepath.Join(dir, "/resources/countries.json"))
	json.Unmarshal(countries, &callingCodes)

	mnos, _ := ioutil.ReadFile(filepath.Join(dir, "/resources/mnos.json"))
	json.Unmarshal(mnos, &mnoIdentifiers)
}

type MsisdnInfo struct {
	CountryCode      string `json:"country_code"`
	DialingNumber    string `json:"dialing_number"`
	MnoIdentifier    string `json:"mno_identifier"`
	SubscriberNumber string `json:"subscriber_name"`
}

func ParseMsisdn(msisdn string) (MsisdnInfo, error) {
	start := findMsisdnStart(msisdn)
	for i := start + 1; i <= len(msisdn); i++ {
		dialingNumber := msisdn[start:i]
		code, ok := callingCodes[dialingNumber]
		if ok {
			subscriberNumber := msisdn[i:]
			return MsisdnInfo{CountryCode: code,
				DialingNumber:    dialingNumber,
				MnoIdentifier:    findMnoIdentifier(code, subscriberNumber),
				SubscriberNumber: subscriberNumber}, nil
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

func findMnoIdentifier(countryCode string, subscriberNumber string) string {
	mnos := mnoIdentifiers[countryCode]
	for i := 1; i <= len(subscriberNumber); i++ {
		var ok bool
		mnoIdentifier, ok := mnos[subscriberNumber[0:i]]
		if ok {
			return mnoIdentifier
		}
	}
	return "unknown"
}
