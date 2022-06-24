package main

import (
	"net/http"
	"strings"
	"fmt"
	"log"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ipAddress := r.RemoteAddr
		fwdAddress := r.Header.Get("X-Forwarded-For") // capitalisation doesn't matter
		if fwdAddress != "" {
			// Got X-Forwarded-For
			ipAddress = fwdAddress // If it's a single IP, then awesome!

			// If we got an array... grab the first IP
			ips := strings.Split(fwdAddress, ", ")
			if len(ips) > 1 {
				ipAddress = ips[0]
			}
		}

        fmt.Fprintf(w, ipAddress)
    })

    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "OK")
    })

    log.Fatal(http.ListenAndServe(":80", nil))

}

/*
func main() {
	// Declare a new router
	r := mux.NewRouter()

	// This is where the router is useful, it allows us to declare methods that
	// this path will be valid for
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/ping", pingHandler).Methods("GET")
	http.Handle("/", r)

	// We can then pass our router (after declaring all our routes) to this method
	// (where previously, we were leaving the second argument as nil)
	http.ListenAndServe(":80", r)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,  "OK")
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	//
	ipAddress := r.RemoteAddr
	fwdAddress := r.Header.Get("X-Forwarded-For") // capitalisation doesn't matter
	if fwdAddress != "" {
		// Got X-Forwarded-For
		ipAddress = fwdAddress // If it's a single IP, then awesome!

		// If we got an array... grab the first IP
		ips := strings.Split(fwdAddress, ", ")
		if len(ips) > 1 {
			ipAddress = ips[0]
		}
	}

	w.Write([]byte(ipAddress))
}
*/