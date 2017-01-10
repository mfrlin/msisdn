package parser

import (
	"testing"
)
// TODO: research unit testing libraries. this is too much boilerplate code

func TestFindMsisdnStart(t *testing.T) {
	var start int
	start = findMsisdnStart("38631123123")
	if start != 0 {
		t.Errorf("wrong msisdn start. got=%d", start)
	}
	start = findMsisdnStart("+38631123123")
	if start != 1 {
		t.Errorf("wrong msisdn start. got=%d", start)
	}
	start = findMsisdnStart("0038631123123")
	if start != 2 {
		t.Errorf("wrong msisdn start. got=%d", start)
	}
}

func TestResources(t *testing.T) {
	if l := len(mno_identifiers); l != 1 {
		t.Errorf("wrong number of keys in mno_identifiers. got=%d", l)
	}

	si_mnos := mno_identifiers["SI"]
	if l := len(si_mnos); l != 9 {
		t.Errorf("wrong number of keys in SI mnos. got=%d", l)

	}

	if l := len(calling_codes); l != 2 {
		t.Errorf("wrong number of keys in calling_codes. got=%d", l)
	}
}

func TestParseMsisdn(t *testing.T) {
	numbers := []string{"38631123123", "+38631123123", "0038631123123"}

	for _, number := range numbers {
		info, _ := ParseMsisdn(number)
		if info.dialing_number != "386" {
			t.Errorf("dialing number wrong. got=%s", info.dialing_number)
		}

		info, err := ParseMsisdn("+38631123123")
		if err != nil {
			t.Errorf("err not nil. got=%s", err)
		}
		if info.country_code != "SI" {
			t.Errorf("country code wrong. got=%s", info.country_code)
		}
		if info.dialing_number != "386" {
			t.Errorf("dialing number wrong. got=%s", info.dialing_number)
		}
		if info.mno_identifier != "Mobitel" {
			t.Errorf("mno identifier wrong. got=%s", info.mno_identifier)
		}
		if info.subscriber_number != "31123123" {
			t.Errorf("subscriber number wrong. got=%s", info.subscriber_number)
		}

		info, err = ParseMsisdn("+38531123123")
		if err != nil {
			t.Errorf("err not nil. got=%s", err)
		}
		if info.country_code != "HR" {
			t.Errorf("country code wrong. got=%s", info.country_code)
		}
		if info.dialing_number != "385" {
			t.Errorf("dialing number wrong. got=%s", info.dialing_number)
		}
		if info.mno_identifier != "unknown" {
			t.Errorf("mno identifier wrong. got=%s", info.mno_identifier)
		}
		if info.subscriber_number != "31123123" {
			t.Errorf("subscriber number wrong. got=%s", info.subscriber_number)
		}

		info, err = ParseMsisdn("+12331123123")
		if err == nil {
			t.Error("err is nil but should not be")
		}
		if err.Error() != "Country code not found." {
			t.Errorf("wrong error message. got=%s", err)

		}


	}


}
