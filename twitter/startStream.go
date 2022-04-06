package twitter

func StartStream(client Client) {
	// Add Filters to Stream
	addFilters(client)

	// Start listening to Stream
	go func(twitterClient Client) {
		listenToStream(client)
	}(client)
}
