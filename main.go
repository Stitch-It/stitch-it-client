package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/caarlos0/env/v6"
)

func main() {
	// Parse env variables for Twitter API keys
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Create a new httpClient Bearer Token
	// for making calls to the Twitter API
	httpClient := &http.Client{}

	image := make(chan string)

	client := Client{
		conf:             cfg,
		http:             httpClient,
		imageUrlAndSizes: image,
	}

	// bytesBuffer := bytes.NewBufferString("testing")

	// decodedBuffer := base64.NewDecoder(base64.StdEncoding, bytesBuffer)

	// io.Copy(os.Stdout, decodedBuffer)

	// Run separate server in goroutine so users can
	// make requests and we can consume the Twitter
	// API at the same time
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			println("hello")
		})

		log.Fatal(http.ListenAndServe(":3030", nil))
	}()

	// Add Filters to Stream
	addFilters(client)

	go func(twitterClient Client) {
		listenToStream(client)
	}(client)

	for imageUrlAndSize := range client.imageUrlAndSizes {
		imgUrlAndSize := imageUrlAndSize

		if imgUrlAndSize == "" {
			continue
		} else {
			imgUrl := strings.Split(imgUrlAndSize, "@-@")[0]
			imgSize := strings.Split(imgUrlAndSize, "@-@")[1]

			fileName, b, _ := downloadImage(imgUrl)

			resizeImage(fileName, b, imgSize)

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
		}
	}
}
