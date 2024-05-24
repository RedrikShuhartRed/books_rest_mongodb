package routes

import (
	"github.com/RedrikShuhartRed/books_rest_mongodb/handlers"
	"github.com/gin-gonic/gin"
)

var RegisterMovieRoutes = func(router *gin.Engine) {

	router.Handle("GET", "/movies", handlers.Get)
}
