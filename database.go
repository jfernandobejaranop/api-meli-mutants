package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = MeliContextDB().Collection("MutantsTransactions")

const uri = "mongodb+srv://root:admin@atlascluster.fewo3js.mongodb.net/?retryWrites=true&w=majority"

func MeliContextDB() *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 1000)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("MELI")
}

func saveTransactions(rst Mutant) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), rst)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
