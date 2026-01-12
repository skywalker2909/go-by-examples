package main

import (
	"fmt"
	/*
	  io package provides functionality for performing various input and
	  output operations/
	*/
	"io"
	/*
		net/http package provides http client and server implementations
		that can be used to work with different http requests and an http
		server to serve http requests.
	*/
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// Exit the process with a status code of 1.
			os.Exit(1)
		}

		// Reads the entire response from the response body stream.
		bytes, err := io.ReadAll(resp.Body)
		/*
			Close the response body after the response is read to avoid
			leaking resources.
		*/
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", bytes)
	}
}
