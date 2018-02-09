package main

import (
	"flag"
	"log"
	"net/http"
)

// var uri = flag.String(name string, value string, usage string)
var flURI = flag.String("uri", "localhost:8765",
	`Supply a string to use as the server uri.
    Example:
       ./server --uri localhost:3456
`)

func main() {
	flag.Parse()
	log.Printf("using uri: %v", *flURI)
	http.HandleFunc("/checkin", checkinHandler)

	//	http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	log.Fatal(http.ListenAndServe(*flURI, nil))
}

func checkinHandler(w http.ResponseWriter, r *http.Request) {

}
