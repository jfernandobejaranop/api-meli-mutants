package database

import (
	"context"
	"log"

	"github.com/api-meli-mutants/mutant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://root:admin@atlascluster.fewo3js.mongodb.net/?retryWrites=true&w=majority"

var collection = MeliContextDB().Collection("MutantsTransactions")

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

func SaveTransactions(rst mutant.MutantInfo) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), rst)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CountDocuments(filter bson.M) int64 {
	cantMutants, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return cantMutants
}

func calculateRatio(num1 int, num2 int) float32 {
	if num1 > 0 && num2 > 0 {
		return float32(num1 / num2)
	} else {
		return 0
	}
}

func StatisticsBD() status {
	cantMutants := int(CountDocuments(bson.M{"ismutant": true}))
	cantNoMutants := int(CountDocuments(bson.M{"ismutant": false}))

	return status{
		CountMutant: cantMutants,
		CountHuman:  cantNoMutants,
		Ratio:       calculateRatio(cantMutants, cantNoMutants),
	}
}
