package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func listenToStream(client Client) {
	req, err := http.NewRequest(http.MethodGet, "https://api.twitter.com/2/tweets/search/stream?tweet.fields=text,attachments,source&expansions=author_id,attachments.media_keys&media.fields=url", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	bearer := fmt.Sprintf("Bearer %s", client.conf.BearerToken)
	req.Header.Set("Authorization", bearer)

	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// var done bool

	reader := bufio.NewReader(resp.Body)
	for {
		bts, _ := reader.ReadBytes('\n')

		handleTweet(bts, client)
	}
}

func addFilters(client Client) {
	f := Filter{
		Value: "@stitchitart has:media",
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
