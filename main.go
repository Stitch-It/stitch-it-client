package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/caarlos0/env/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Create values to be storeed in out Client
	// struct, including our MongoClient and
	// our context
	tweet := make(chan Tweet)

	ctx, mongoClient := MongoConnect(cfg.MongoUri)
	mongoClient.Collection = mongoClient.ConfigureCollection(cfg.DatabaseName, cfg.CollectionName)

	client := Client{
		conf:         cfg,
		http:         httpClient,
		imageTweet:   tweet,
		mongoContext: ctx,
		mongoClient:  mongoClient,
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

	// Add Filters to Stream
	addFilters(client)

	// Start listening to Stream
	go func(twitterClient Client) {
		listenToStream(client)
	}(client)

	// Listen on client.imgTweet channel
	// for a Tweet to process
	for imgTweet := range client.imageTweet {
		if imgTweet.Next {
			continue
		} else {
			imgUrl := imgTweet.MediaUrl
			imgSize := imgTweet.Text
			// tweetId := imgTweet.Id

			fileName, b, _ := downloadImage(imgUrl)

			resizeImage(fileName, b, imgSize)

			// After downloading and resizing the
			// image, create a MongoDoc with image
			// metadata to beinserted into the
			// database
			imgDoc := MongoDoc{
				ImageName: fileName,
				ImagePath: "/images/" + fileName,
			}

			// Insert the image metadata into
			// the database
			res, err := client.mongoClient.InsertImageMetaData(client.mongoContext, imgDoc)
			if err != nil {
				println(fmt.Sprintf("Insert Doc error: %v", err))
			}
			splitOne := strings.Split(res.InsertedID.(primitive.ObjectID).String(), "(")[1]
			splitTwo := strings.Split(splitOne, ")")[0]
			imgId := splitTwo[1 : len(splitTwo)-1]
			if imgId != "" {
				println(imgId)
			}

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
		}
	}
}
