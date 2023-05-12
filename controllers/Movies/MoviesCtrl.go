package Movies

import (
	"awesomeProject/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

var movies = []entities.Movie{
	{ID: "1", Title: "Blue Train", Director: "John Coltrane", Length: 56.99},
	{ID: "2", Title: "Jeru", Director: "Gerry Mulligan", Length: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Director: "Sarah Vaughan", Length: 39.99},
}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	for _, a := range movies {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}

func CreateMovies(c *gin.Context) {
	var newMovie entities.Movie

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	// Add the new album to the slice.
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}
