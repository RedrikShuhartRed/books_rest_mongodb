package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/RedrikShuhartRed/books_rest_mongodb/db"
	"github.com/RedrikShuhartRed/books_rest_mongodb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll(c *gin.Context) {
	ctx := context.Background()
	client := db.GetDB()
	collection := client.Database("moviebox").Collection("movies")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Ошибка при получении данных из коллекции:", err)
		return
	}

	defer cursor.Close(ctx)

	var movies []models.Movie
	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			continue
		}
		if movie.Title != "" {
			movies = append(movies, movie)
		}
	}

	c.JSON(http.StatusOK, movies)
}

func GetMoviesByDirector(c *gin.Context) {
	ctx := context.Background()
	client := db.GetDB()
	collection := client.Database("moviebox").Collection("movies")
	director := c.Param("director")

	cursor, err := collection.Find(ctx, bson.M{"director": director})
	if err != nil {
		log.Println("Ошибка при получении данных из коллекции:", err)
		return
	}

	var movies []models.Movie
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Ошибка при декодировании документа:", err)
			continue
		}

		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Ошибка при чтении данных из коллекции:", err)
		return
	}

	c.JSON(http.StatusOK, movies)
}

func AddMovies(c *gin.Context) {
	ctx := context.Background()
	client := db.GetDB()
	collection := client.Database("moviebox").Collection("movies")
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Ошибка чтения данных тела запроса AddMovies: %s\n", err)
	}

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
	c.JSON(http.StatusOK, movies)

}

func SortByRating(c *gin.Context) {
	ctx := context.Background()
	client := db.GetDB()
	collection := client.Database("moviebox").Collection("movies")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "rating", Value: -1}})
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Println("Ошибка при получении данных из коллекции:", err)
		return
	}

	var movies []models.Movie
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Ошибка при декодировании документа:", err)
			continue
		}

		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Ошибка при чтении данных из коллекции:", err)
		return
	}

	c.JSON(http.StatusOK, movies)
}
