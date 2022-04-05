package main

func handleTweet(bytes []byte, client Client) bool {
	var done bool = false

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
			urlAndSize := tweet.MediaUrl + "@-@" + tweet.Text

			client.imageUrlAndSizes <- urlAndSize
			client.imageUrlAndSizes <- ""

			done = true
		}
	}

	return done
}
