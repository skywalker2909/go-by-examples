package main

import (
	"fmt"
	/*
		log package provides functions for logging information in the
		form of errors, warnings, debug and info etc.
	*/
	"log"
	/*
		net/http package provides http client and server implementations
		that can be used to work with different http requests and an http
		server to serve http requests.
	*/
	"net/http"
)

func main() {
	/*
		Register an http request handler function for handling
		all incoming requests that begin with a '/' in the URL.

		For example: /. /help, /about-me etc. in
		http://localhost:8000/
		http://localhost:8000/help
		http://localhost:8000/about-me
	*/
	http.HandleFunc("/", handler)

	/*
		'http.ListenAndServe' starts an http server and listens for
		incoming requests on port number 8000

		It returns an error on failure to listen and serve which is
		passed on to 'log.Fatal' as a parameter to log and exit.
	*/
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
