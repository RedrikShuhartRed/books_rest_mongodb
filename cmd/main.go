package main

import (
	"log"
	"net/http"

	"github.com/RedrikShuhartRed/books_rest_mongodb/db"
	"github.com/RedrikShuhartRed/books_rest_mongodb/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	client := db.ConnectDb()
	r := gin.Default()

	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)

	r.Run(":8080")

	db.CloseDb(client)

}
