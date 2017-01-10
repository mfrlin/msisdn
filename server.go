package main

import (
	"msisdn/parser"
	"net/http"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msisdn := r.URL.Path[1:]
	info, err := parser.ParseMsisdn(msisdn)
	out := json.NewEncoder(w)
	if err != nil {
		out.Encode(err.Error())
		return
	}
	out.Encode(info)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
