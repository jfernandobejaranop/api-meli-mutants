package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func SaveTransactions(rst Mutant) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), rst)
	if err != nil {
		log.Fatal(err)
	}
	Stats()
	return result
}

func Stats() status {
	var resultsRatio int64

	filterTrue := bson.M{"ismutant": true}
	filterFalse := bson.M{"ismutant": false}
	resulttrue, err := collection.CountDocuments(context.TODO(), filterTrue)
	resultFalse, err := collection.CountDocuments(context.TODO(), filterFalse)

	if resulttrue > 0 && resultFalse > 0 {
		resultsRatio = resulttrue / resultFalse
	} else {
		resultsRatio = 0
	}

	var rst = status{
		CountMutant: resulttrue,
		CountHuman:  resultFalse,
		Ratio:       resultsRatio,
	}

	b, err := json.Marshal(rst)

	fmt.Println("Resultado count : ", string(b), err)

	return rst
}

type status struct {
	CountMutant int64 `json:"count_mutant_dna"`
	CountHuman  int64 `json:"count_human_dna"`
	Ratio       int64 `json:"ratio"`
}
