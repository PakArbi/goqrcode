package goqrcode

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func insertPayload() {
	// Establish a connection to your MongoDB instance
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		panic(err)
	}

	// Assign the collection to your global variable
	collection = client.Database("Pakarbi").Collection("codeqr")
}

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("ulbi").Collection("users")
}

// InsertPayload inserts a user into MongoDB
func InsertPayload(payload Payload) error {
	_, err := collection.InsertOne(context.Background(), payload)
	return err
}