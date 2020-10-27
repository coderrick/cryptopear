package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name   string
	Prices []string
}

func products(w http.ResponseWriter, req *http.Request) {

	profile := Product{"DeepLee Projector", []string{"71.00", "89.00"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", products)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
