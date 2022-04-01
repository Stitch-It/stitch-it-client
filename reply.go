package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func reply(client *twitter.Client, tweetId int64, userName string) {

	replyStatus := fmt.Sprintf("@%v Thanks for mentioning!!", userName)

	params := &twitter.StatusUpdateParams{
		Status:            replyStatus,
		InReplyToStatusID: tweetId,
	}

	_, _, err := client.Statuses.Update(replyStatus, params)
	if err != nil {
		log.Fatal(err)
	}

}
