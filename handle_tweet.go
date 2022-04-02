package main

import "fmt"

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

				// Generate Pattern

				// Reply with Pattern

				// Signal done
				done <- true
			}
		}
	}(done, tweet)
}
