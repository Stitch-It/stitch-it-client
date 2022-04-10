package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReplyUrl(client Client, tweet Tweet, fileName string) {
	replyId := ReplyId{
		InReplyToStatusId: tweet.Id,
	}

	var reply ReplyTweet

	reply.Text = "hello"
	reply.ReplyToId = replyId

	twt := make(map[string]interface{})
	twt["text"] = reply.Text
	twt["reply"] = reply.ReplyToId

	body, _ := json.Marshal(twt)

	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/2/tweets", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := client.Oauth.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	println(string(reader))
}

func ReplyError(client Client, tweet Tweet, fileName string, replyError string) {
	replyId := ReplyId{
		InReplyToStatusId: tweet.Id,
	}

	var reply ReplyTweet

	reply.Text = replyError
	reply.ReplyToId = replyId

	twt := make(map[string]interface{})
	twt["text"] = reply.Text
	twt["reply"] = reply.ReplyToId

	body, _ := json.Marshal(twt)

	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/2/tweets", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := client.Oauth.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	println(string(reader))
}
