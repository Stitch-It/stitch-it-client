package main

import (
	"fmt"
	"strconv"

	"github.com/gofrs/flock"
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
			storeImageInServer(tweet, client)
		}
	}
}

func storeImageInServer(tweet Tweet, client Client) {
	quit := make(chan bool)
	go func(quit chan bool, tweet Tweet) {
		for {
			select {
			case <-quit:
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

				fileLock := flock.New(fileName)

				println(strconv.FormatBool(fileLock.Locked()))

				// Signal done
				quit <- true
			}
		}
	}(quit, tweet)
}
