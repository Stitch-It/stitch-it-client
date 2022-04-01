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
	req, err := http.NewRequest(http.MethodPost, "", &buf)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Change this to for loop consuming string
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(string(bytes))
}
