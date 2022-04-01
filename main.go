package main

import (
	"fmt"
	"net/http"

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

	// Add Filters to Stream
	addFilters(client)

	// Begin listening to stream with httpClient
	listenToStream(client)
}
