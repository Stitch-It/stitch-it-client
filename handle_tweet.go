package main

import (
	"fmt"
)

func handleTweet(bytes []byte, client Client) {
	// This check handles sporadic empty messages
	if len(bytes) >= 0 {
		tweet := Tweet{
			Error: false,
		}
		tweet = extractValues(bytes, tweet)

		// Check for empty tweet.MediaUrl to
		// prevent crash from panic in processing
		// images
		if tweet.MediaUrl != "" {
			createGoRoutineForTweet(tweet, client)
		}
	}
}

func createGoRoutineForTweet(tweet Tweet, client Client) {
	done := make(chan bool)
	go func(done chan bool, tweet Tweet) {
		for {
			select {
			case <-done:
				return
			default:
				// Download Image
				fileName, b, err := downloadImage(tweet.MediaUrl, tweet.AuthorName)
				if err != nil {
					fmt.Printf("%v\n", err)
				}

				// Resize the image
				resizeImage(fileName, b, tweet.Text)

				sendProcessedImageToServer(fileName, client)

				// Reply to tweet with URL to download
				// Excel pattern

				// Signal done
				done <- true
			}
		}
	}(done, tweet)
}
