package routes

import (
	"github.com/RedrikShuhartRed/books_rest_mongodb/handlers"
	"github.com/gin-gonic/gin"
)

var RegisterMovieRoutes = func(router *gin.Engine) {

	router.GET("/movies", handlers.GetAll)
	router.GET("/movies/:director", handlers.GetMoviesByDirector)
	router.POST("/movies/addmovies", handlers.AddMovies)
	router.GET("/movies/sort/rating", handlers.SortByRating)
}
