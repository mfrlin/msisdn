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
	if l := len(mnoIdentifiers); l != 1 {
		t.Errorf("wrong number of keys in mnoIdentifiers. got=%d", l)
	}

	siMnos := mnoIdentifiers["SI"]
	if l := len(siMnos); l != 9 {
		t.Errorf("wrong number of keys in SI mnos. got=%d", l)

	}

	if l := len(callingCodes); l != 2 {
		t.Errorf("wrong number of keys in callingCodes. got=%d", l)
	}
}

func TestParseMsisdn(t *testing.T) {
	numbers := []string{"38631123123", "+38631123123", "0038631123123"}

	for _, number := range numbers {
		info, _ := ParseMsisdn(number)
		if info.DialingNumber != "386" {
			t.Errorf("dialing number wrong. got=%s", info.DialingNumber)
		}

		info, err := ParseMsisdn("+38631123123")
		if err != nil {
			t.Errorf("err not nil. got=%s", err)
		}
		if info.CountryCode != "SI" {
			t.Errorf("country code wrong. got=%s", info.CountryCode)
		}
		if info.DialingNumber != "386" {
			t.Errorf("dialing number wrong. got=%s", info.DialingNumber)
		}
		if info.MnoIdentifier != "Mobitel" {
			t.Errorf("mno identifier wrong. got=%s", info.MnoIdentifier)
		}
		if info.SubscriberNumber != "31123123" {
			t.Errorf("subscriber number wrong. got=%s", info.SubscriberNumber)
		}

		info, err = ParseMsisdn("+38531123123")
		if err != nil {
			t.Errorf("err not nil. got=%s", err)
		}
		if info.CountryCode != "HR" {
			t.Errorf("country code wrong. got=%s", info.CountryCode)
		}
		if info.DialingNumber != "385" {
			t.Errorf("dialing number wrong. got=%s", info.DialingNumber)
		}
		if info.MnoIdentifier != "unknown" {
			t.Errorf("mno identifier wrong. got=%s", info.MnoIdentifier)
		}
		if info.SubscriberNumber != "31123123" {
			t.Errorf("subscriber number wrong. got=%s", info.SubscriberNumber)
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
