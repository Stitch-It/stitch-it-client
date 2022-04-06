package main

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	conf         *Config
	http         *http.Client
	imageTweet   chan Tweet
	mongoContext context.Context
	mongoClient  MongoClient
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
	//Next             bool
}

type MongoClient struct {
	Options    *options.ClientOptions
	Mongo      *mongo.Client
	Collection *mongo.Collection
}

type MongoDoc struct {
	ImageName string `json:"ImageName Str"`
	ImagePath string `json:"ImagePath Str"`
}
