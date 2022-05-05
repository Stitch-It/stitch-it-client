package main

import (
	"fmt"
	"net/http"

	gen "github.com/Stitch-It/stitch-it/generate-pattern"
	imgHdl "github.com/Stitch-It/stitch-it/image-process"
	"github.com/Stitch-It/stitch-it/twitter"
	"github.com/caarlos0/env/v6"
	"github.com/dghubble/oauth1"
)

func main() {

	// Lines 19 through 48 are relevant to this software only in its Twitter bot iteration
	// It can largely be ignored for the purpose of refactoring to use the Fyne framework

	// Parse env variables for Twitter credentials
	cfg := &twitter.Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Create a new httpClient that will use
	// the Bearer Token to authorize into
	// endpoint for listening to Tweet stream
	httpClient := &http.Client{}

	// Create oauthClient for authorization into
	// endpoint for replying to tweets
	config := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)

	oauthClient := config.Client(oauth1.NoContext, token)

	// tweet is a channel for the stream to send
	// individual Tweets on for processing, pattern
	// generation, and Tweet response in separate
	// goroutines for each tweet
	tweet := make(chan twitter.Tweet)

	client := twitter.Client{
		Conf:       cfg,
		Http:       httpClient,
		Oauth:      oauthClient,
		ImageTweet: tweet,
	}

	// Run separate server in goroutine so users can
	// make requests (responding to tweets) and we can consume the Twitter
	// API at the same time
	//go func() {
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		println("hello")
	//	})
	//
	//	log.Fatal(http.ListenAndServe(":3030", nil))
	//}()

	// Here is where the stream for listing for individual tweets is started
	twitter.StartStream(client)

	// Listen on client.ImageTweet channel
	// for a Tweet to process
	for imgTweet := range client.ImageTweet {
		done := make(chan bool)
		go func(twt twitter.Tweet) {
			for {
				select {
				case <-done:
					return
				default:
					imgUrl := twt.MediaUrl
					imgSize := twt.Text

					// This is where the main functionality that will be ported over to using the Fyne framework will be
					fileName, b, _ := imgHdl.DownloadImage(imgUrl)

					imgHdl.ResizeImage(fileName, b, imgSize)

					gen.GenerateExcelPattern(fileName, twt.AuthorScreenName)

					twitter.ReplyUrl(client, twt, fileName)

					done <- true
				}
			}

		}(imgTweet)
	}
}
