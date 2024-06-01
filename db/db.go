package db

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/RedrikShuhartRed/books_rest_mongodb/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientDb *mongo.Client

func ConnectDb() *mongo.Client {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Printf("Ошибка подключения к БД: %s\n", err)
	}

	data, err := os.ReadFile("D:/project/rest_api_mongodb/final-db.json")
	if err != nil {
		log.Printf("Ошибка чтения файла: %s\n", err)
	}

	db := client.Database("moviebox")
	collection := db.Collection("movies")
	var movies []models.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		log.Printf("Ошибка Unmarshal при создании коллекции: %s\n", err)
	}

	for _, film := range movies {
		_, err := collection.InsertOne(ctx, film)
		if err != nil {
			log.Printf("Ошибка InsertOne при создании коллекции: %s\n", err)
		}
	}

	clientDb = client
	return clientDb

}

func GetDB() *mongo.Client {
	return clientDb
}

func CloseDb(client *mongo.Client) {
	ctx := context.Background()
	client.Disconnect(ctx)
}
