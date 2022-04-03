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

	// // Add Filters to Stream
	addFilters(client)

	// var proc *os.Process

	// // Begin listening to stream with httpClient
	listenToStream(client)
}
