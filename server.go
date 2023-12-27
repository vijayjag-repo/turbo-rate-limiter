package main

import (
	"fmt"
	"net/http"
)

func limitedHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Limited, don't over use me!\n")
}

func unlimitedHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Unlimited, LFG!\n")
}

func main() {
	http.HandleFunc("/limited", limitedHandler)
	http.HandleFunc("/unlimited", unlimitedHandler)

	//  curl http://127.0.0.1:8090/limited
	http.ListenAndServe(":8090", nil)
}
