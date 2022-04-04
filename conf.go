package main

type Config struct {
	BearerToken    string `env:"BEARER_TOKEN"`
	MongoUri       string `env:"MONGO_URI"`
	DatabaseName   string `env:"DATABASE_NAME"`
	CollectionName string `env:"COLLECTION_NAME"`
}
