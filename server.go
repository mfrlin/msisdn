package main

import (
	"encoding/json"
	"flag"
	"msisdn/parser"
	"net/http"
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
	var port = flag.String("port", "8080", "Port on which server should listen")
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*port, nil)
}
