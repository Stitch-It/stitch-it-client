package main

type Config struct {
	ConsumerKey    string `env:"CONSUMER_KEY"`
	ConsumerSecret string `env:"CONSUMER_SECRET"`
	AccessToken    string `env:"ACCESS_TOKEN"`
	AccessSecret   string `env:"ACCESS_SECRET"`
}
