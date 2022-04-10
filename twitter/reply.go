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

	reply := ReplyTweet{
		Text:    "hello",
		ReplyId: replyId,
	}

	body, _ := json.Marshal(reply)

	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/2/tweets", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	bearer := fmt.Sprintf("Bearer %s", client.Conf.BearerToken)
	req.Header.Set("Authorization", bearer)

	resp, err := client.Http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	println(string(reader))
}

func ReplyError(client Client, tweet Tweet, fileName string) {

	// replyStatus := fmt.Sprintf("@%v Thanks for mentioning!!", userName)

}
