package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (*mongo.Collection) {


	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_URL"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return client.Database("Guardian_Users").Collection("users")
}
