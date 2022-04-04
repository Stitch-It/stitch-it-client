package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/caarlos0/env/v6"
)

type Client struct {
	conf *Config
	http *http.Client
}

func main() {
	// Parse env variables for Twitter API keys
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Create a new httpClient Bearer Token
	// for making calls to the Twitter API
	httpClient := &http.Client{}

	client := Client{
		conf: cfg,
		http: httpClient,
	}

	bytesBuffer := bytes.NewBufferString("testing")

	decodedBuffer := base64.NewDecoder(base64.StdEncoding, bytesBuffer)

	io.Copy(os.Stdout, decodedBuffer)

	// Run separate server in goroutine so users can
	// make requests and we can consume the Twitter
	// API at the same time
	// go func() {
	// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		println("hello")
	// 	})

	// 	log.Fatal(http.ListenAndServe(":3030", nil))
	// }()

	// // Add Filters to Stream
	addFilters(client)

	// var tweetProcessed bool

	// for !tweetProcessed {
	// 	tweetProcessed = listenToStream(client)

	// 	continue
	// }

	// Begin listening to stream with httpClient
	listenToStream(client)
}
