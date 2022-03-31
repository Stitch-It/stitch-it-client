package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// Parse env variables for Twitter API keys
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v", err)
	}

	// Configure a new httpClient to pass to twitter.NewClient()
	conf := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	httpClient := conf.Client(oauth1.NoContext, token)

	// Create new Twitter client
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account @%s (%s)\n", user.ScreenName, user.Name)

	// Convenience demultiplexer to type switch messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		// processing picture and replying with
		// pattern will be handled here
		println(tweet.Text)
		println(fmt.Sprintf("@%s", tweet.User.ScreenName))
		// the line below downloads the image
		// will be replaced with
		// tweet.Entities.Media[0].MediaURLHttps
		// to save image from tweet
		// downloadImage("https://github.com/syke99/go-c2dmc/blob/main/img/Screenshot%202021-12-04%20101419.png?raw=true")
	}

	// Filter stream
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"@ArtStitchit"}, // follow tweets mentioning this user
		StallWarnings: twitter.Bool(true),       // include a Stall Warning
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	go demux.HandleChan(stream.Messages)

	// Wait gor SIGINT and SIGTERM (Hitting CTRL-C)
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	stream.Stop()
}
