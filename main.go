package main

import (
	"awesomeProject/controllers/Movies"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.

// albums slice to seed record album data.

func main() {

	router := gin.Default()
	router.GET("/movies", Movies.GetMovies)
	router.GET("/movies/:id", Movies.GetMovie)
	router.POST("/movies", Movies.CreateMovies)

	router.Run("localhost:8080")
}
