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

	reader := bufio.NewReader(resp.Body)
	// This will be changed. This is where each
	// individual Tweet will be handled
	for {
		bytes, _ := reader.ReadBytes('\n')

		// This check handles sporadic empty messages
		if len(bytes) > 0 {
			tweet := Tweet{
				Error: false,
			}
			tweet = extractValues(bytes, tweet)

			done := make(chan bool)
			go func(done chan bool, tweet Tweet) {
				for {
					select {
					case <-done:
						return
					default:
						// Download Image
						fileName, err := downloadImage(tweet.MediaUrl, tweet.AuthorName)
						if err != nil {
							fmt.Printf("%v\n", err)
						}

						// Resize the image
						err = processImage(fileName, tweet.Text)
						if err != nil {
							fmt.Printf("%v\n", err)
						}

						// Generate Pattern

						// Reply with Pattern

						// Signal done
						done <- true
					}
				}
			}(done, tweet)
		}
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
