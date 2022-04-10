package twitter

import (
	"net/http"
)

type Config struct {
	BearerToken string `env:"BEARER_TOKEN"`
}

type Client struct {
	Conf       *Config
	Http       *http.Client
	ImageTweet chan Tweet
}

type Filter struct {
	Value string `json:"value"`
}

type FilterRules struct {
	Add []Filter `json:"add"`
}

type Tweet struct {
	Id               string
	AuthorId         string
	AuthorName       string
	AuthorScreenName string
	Text             string
	MediaUrl         string
	Error            bool
}

type ReplyId struct {
	InReplyToStatusId string `json:"in_reply_to_status_id"`
}

type ReplyTweet struct {
	Text string `json:"text"`
	ReplyId
}
