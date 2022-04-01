package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
)

func listenToStream(client *twitter.Client) {

	// Convenience demultiplexer to type switch messages
	demux := twitter.NewSwitchDemux()

	demux.Tweet = func(tweet *twitter.Tweet) {
		// fileName, err := downloadImage(fmt.Sprint(tweet.Entities.Media[0].MediaURLHttps), tweet.User.Name)
		// if err != nil {
		// 	fmt.Printf("%v\n", err)
		// }

		// err = processImage(fileName, tweet.Text)
		// if err != nil {
		// 	fmt.Printf("%v\n", err)
		// }

		reply(client, tweet.ID, tweet.User.ScreenName)
	}

	// Filter stream
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"@StitchItArt"}, // follow tweets mentioning this user
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
