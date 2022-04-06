package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	imgHdl "github.com/syke99/stitch-it/image-process"
	"github.com/syke99/stitch-it/twitter"
)

func main() {
	// Parse env variables for Twitter API keys
	cfg := &twitter.Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Create a new httpClient Bearer Token
	// for making calls to the Twitter API
	httpClient := &http.Client{}

	// Create values to be storeed in out Client
	// struct, including our MongoClient and
	// our context
	tweet := make(chan twitter.Tweet)

	client := twitter.Client{
		Conf:       cfg,
		Http:       httpClient,
		ImageTweet: tweet,
	}

	// Run separate server in goroutine so users can
	// make requests and we can consume the Twitter
	// API at the same time
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			println("hello")
		})

		log.Fatal(http.ListenAndServe(":3030", nil))
	}()

	// // Add Filters to Stream
	// addFilters(client)

	// // Start listening to Stream
	// go func(twitterClient Client) {
	// 	listenToStream(client)
	// }(client)

	twitter.StartStream(client)

	// Listen on client.imgTweet channel
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
					//tweetId := imgTweet.Id

					fileName, b, _ := imgHdl.DownloadImage(imgUrl)

					imgHdl.ResizeImage(fileName, b, imgSize)
					//------------------------------

					// Here is where the reply to the user
					// with the URL for downloading their
					// pattern will go

					// This was just proof of concept that
					// this method works for downloading
					// images and them being able to be deleted
					// as well so that I can age and delete them
					// later
					// time.Sleep(15 * time.Second)

					// err := os.Remove("./images/" + fileName)
					// if err != nil {
					// 	fmt.Printf("%v\n", err)
					// }
					done <- true
				}
			}

		}(imgTweet)
	}
}
