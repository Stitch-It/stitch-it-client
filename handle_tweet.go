package main

import (
	"fmt"
	"sync"
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

			var wg sync.WaitGroup

			wg.Add(1)

			go func(tweet Tweet, client Client) {
				storeImageInServer(tweet, client)
			}(tweet, client)

			wg.Wait()

			println("I got out of the goroutine")
		}
	}
}

func storeImageInServer(tweet Tweet, client Client) {
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

	// fileLock := flock.New(fileName)

	// println(strconv.FormatBool(fileLock.Locked()))

}
