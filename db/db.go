package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RedrikShuhartRed/books_rest_mongodb/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbs *mongo.Database

func ConnectDb() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := os.ReadFile("D:/project/rest_api_mongodb/final-db.json")
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("moviebox")
	collection := db.Collection("movies")
	var movies []models.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		log.Fatal(err)
	}

	for _, film := range movies {
		_, err := collection.InsertOne(ctx, film)
		if err != nil {
			log.Fatal(err)
		}
	}
	dbs = db
	fmt.Println("Данные добавлены в коллекцию")

}

func GetDB() *mongo.Database {
	return dbs
}
