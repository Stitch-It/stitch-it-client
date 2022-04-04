package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(uri string) (context.Context, MongoClient) {
	ctx := context.Background()

	var client MongoClient

	clientOptions := options.Client().ApplyURI(uri)
	println(fmt.Sprintf("ClientOptions Type: %v", reflect.TypeOf(clientOptions)))

	clt, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println(fmt.Sprintf("Mongo.connect() error: %v", err))
		os.Exit(1)
	}
	println(fmt.Sprintf("Client Type: %v", reflect.TypeOf(clt)))

	client.Options = clientOptions
	client.Mongo = clt

	return ctx, client
}

func (client *MongoClient) ConfigureCollection(databaseName string, collectionName string) *mongo.Collection {
	return client.Mongo.Database(databaseName).Collection(collectionName)
}

func (client *MongoClient) InsertImageMetaData(ctx context.Context, doc MongoDoc) (*mongo.InsertOneResult, error) {
	return client.Collection.InsertOne(ctx, doc)
}
