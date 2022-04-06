package twitter

import (
	"net/http"
)

type Config struct {
	BearerToken string `env:"BEARER_TOKEN"`
	// MongoUri       string `env:"MONGO_URI"`
	// DatabaseName   string `env:"DATABASE_NAME"`
	// CollectionName string `env:"COLLECTION_NAME"`
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
