package handlers

import (
	"context"
	"log"

	"github.com/RedrikShuhartRed/books_rest_mongodb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Get gin.HandlerFunc

func GetAll(client *mongo.Client) ([]models.Movie, error) {
	ctx := context.Background()
	db := client.Database("moviebox")
	collection := db.Collection("movies")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Ошибка при подключении к коллекции", err)
		return nil, err
	}
	var movie models.Movie
	var movies []models.Movie
	for cursor.Next(ctx) {

		err := cursor.Decode(&movie)
		if err != nil {
			log.Println("Ошибка при декодировании документа:", err)
			continue
		}

		// Добавляем только если movie не пустой
		if movie.Title != "" {
			movies = append(movies, movie)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	if err := cursor.Close(ctx); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(movies)

	// Get = func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"movies": movies,
	// 	})
	// }
	return movies, nil
}
