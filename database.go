package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MutantsTransactions_cll = MeliContextDB().Collection("MutantsTransactions")

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

func saveTransactions(req Mutant) *mongo.InsertOneResult {

	result, err := MutantsTransactions_cll.InsertOne(context.TODO(), req)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
