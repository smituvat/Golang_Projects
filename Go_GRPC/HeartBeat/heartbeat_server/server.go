package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	// By default grpc listens to 50051

	lis, err := net.Listen("tcp", ":50051")
	handleError(err)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)

	fmt.Println("MongoDB connected")

	err = client.Connect(context.TODO())
	handleError(err)

	collection = client.Database("heartbeat").Collection("heartbeat")

	fmt.Println(lis)

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
