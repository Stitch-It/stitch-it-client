package main

import (
	"net/http"
)

type Client struct {
	conf       *Config
	http       *http.Client
	imageTweet chan Tweet
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
