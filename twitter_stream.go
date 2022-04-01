package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Filter struct {
	Value string `json:"value"`
}

type FilterRules struct {
	Add []Filter `json:"add"`
}

func listenToStream(client Client) {
	req, err := http.NewRequest(http.MethodGet, "https://api.twitter.com/2/tweets/search/stream", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	bearer := fmt.Sprintf("Bearer %s", client.conf.BearerToken)
	req.Header.Set("Authorization", bearer)

	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// This will be changed. This is where each
	// individual Tweet will be read
	for {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		fmt.Println(string(bytes))
	}
}

func addFilters(client Client) {
	f := Filter{
		Value: "@:stitchit has:media",
	}

	filters := []Filter{f}

	rules := FilterRules{
		Add: filters,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(rules)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Add URL for making request to begin listening to stream
	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/2/tweets/search/stream/rules", &buf)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	bearer := fmt.Sprintf("Bearer %s", client.conf.BearerToken)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", bearer)

	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(string(bytes))
}
