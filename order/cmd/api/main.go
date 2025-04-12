package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = "80"
	mongoUrl = "mongodb://mongo:27017"
)

var client *mongo.Client

type Config struct {
}

func main() {
	// connect mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

}

// connectToMongo
func connectToMongo() (*mongo.Client, error) {
	// create conn opts
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// conn
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to mongo: ", err)
		return nil, err
	}

	return c, nil
}
