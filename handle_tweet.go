package main

import (
	"fmt"
)

func handleTweet(bytes []byte) {
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
			createGoRoutineForTweet(tweet)
		}
	}
}

func createGoRoutineForTweet(tweet Tweet) {
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

				// println(fileName)

				// Resize the image
				resizeImage(fileName, tweet.Text)
				if err != nil {
					fmt.Printf("%v\n", err)
				}

				// Instead of generating the excel file
				// here, we'll send the resized image
				// to the Stitch-It file server here,
				// store it there, and dynamically create
				// the pattern when a user makes a GET request
				// to download the pattern
				//
				// sendProcessedImageToServer()

				// Reply to tweet with URL to download
				// Excel pattern

				// Signal done
				done <- true
			}
		}
	}(done, tweet)
}
