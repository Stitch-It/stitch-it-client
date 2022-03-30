package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// Parse env variables for Twitter API keys
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%v", err)
	}

	// Configure a new httpClient to pass to twitter.NewClient()
	conf := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)
	httpClient := conf.Client(oauth1.NoContext, token)

	// Create new Twitter client
	client := twitter.NewClient(httpClient)

	// Convenience demultiplexer to type switch messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		println(tweet.Text)
	}

	// Filter stream
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"@stitch-it"}, // follow tweets mentioning this user
		StallWarnings: twitter.Bool(true),     // include a Stall Warning
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		fmt.Printf("%v", err)
	}

	go demux.HandleChan(stream.Messages)

	// Wait gor SIGINT and SIGTERM (Hitting CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	stream.Stop()
}
