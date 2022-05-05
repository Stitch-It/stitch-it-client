package twitter

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
			// tweet.Next = false

			client.ImageTweet <- tweet

			// nextTweet := Tweet{
			// 	Next: true,
			// }

			// client.imageTweet <- nextTweet
		}
	}
}
