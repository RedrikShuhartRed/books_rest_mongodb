package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/RedrikShuhartRed/books_rest_mongodb/db"
	"github.com/RedrikShuhartRed/books_rest_mongodb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(c *gin.Context) {
	ctx := context.Background()
	client := db.GetDB()
	collection := client.Database("moviebox").Collection("movies")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подключении к коллекции"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных из коллекции"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении данных из коллекции"})
		log.Println("Ошибка при чтении данных из коллекции:", err)
		return
	}

	c.JSON(http.StatusOK, movies)
}
