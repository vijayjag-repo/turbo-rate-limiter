package main

import (
	"fmt"
	"net/http"
	"turbo-rate-limiter/strategy"
)

func limitedHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Limited, don't over use me!\n")
}

func unlimitedHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Unlimited, LFG!\n")
}

func main() {
	//  curl http://127.0.0.1:8090/limited
	http.HandleFunc("/limited", limitedHandler)
	//  curl http://127.0.0.1:8090/unlimited
	http.HandleFunc("/unlimited", unlimitedHandler)

	// Configure simple token bucket
	strategy.ConfigureTokenBucket(20, 30)

	http.ListenAndServe(":8090", nil)
}
