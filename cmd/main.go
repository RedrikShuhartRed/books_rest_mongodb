package main

import (
	"net/http"

	"github.com/RedrikShuhartRed/books_rest_mongodb/db"
	"github.com/RedrikShuhartRed/books_rest_mongodb/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDb()
	r := gin.Default()
	//_, client := db.GetDB()
	// movies, _ := handlers.GetAll(client)
	//routes.RegisterMovieRoutes(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/movies", handlers.GetAll)
	r.GET("/movies/:director", handlers.GetMoviesByDirector)
	r.Run(":8080")

}
