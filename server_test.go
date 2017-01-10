package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/+38631123123", nil)
	w := httptest.NewRecorder()
	http.HandlerFunc(handler).ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("return status should be 200. got=%v", w.Code)
	}
	var info map[string]string
	json.Unmarshal(w.Body.Bytes(), &info)
	if info["country_code"] != "SI" {
		t.Errorf("country_code wrong. got=%s", info["country_code"])
	}
	if info["dialing_number"] != "386" {
		t.Errorf("dialing_number wrong. got=%s", info["dialing_number"])
	}
	if info["subscriber_number"] != "31123123" {
		t.Errorf("subscriber_number wrong. got=%s", info["subscriber_number"])
	}
	if info["mno_identifier"] != "Mobitel" {
		t.Errorf("mno_identifier wrong. got=%s", info["mno_identifier"])
	}
}
