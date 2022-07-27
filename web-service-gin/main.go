package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// cover represents the structure of the data of each cover
type cover struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	Artist           string    `json:"artist"`
	PremierMonthYear float64   `json:"premier"`
	TimeAdded        time.Time `json:"time_added"`
}

// slice which contains the data of each album
var covers = []cover{
	{ID: "1", Title: "Interstellar", Artist: "Hans Zimmer", PremierMonthYear: 10.2014},
	{ID: "2", Title: "Building Light", Artist: "Anne Sophie Versnaeyen", PremierMonthYear: 2018},
	{ID: "3", Title: "Zen", Artist: "Torin Borrowdale", PremierMonthYear: 03.2018},
	{ID: "4", Title: "Married Life", Artist: "Jason Lyle Black", PremierMonthYear: 10.2015},
}

func main() {
	router := gin.Default()
	router.GET("/covers", getCovers)
	router.GET("/covers/:id", getCoverByID)
	router.POST("/covers", postCovers)

	router.Run("localhost:8080")
}

// gin.Context carries request details, validates and serializes JSON, and more
func getCovers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, covers)
}

func postCovers(c *gin.Context) {
	var newCover cover

	if err := c.BindJSON(&newCover); err != nil {
		return
	}

	covers = append(covers, newCover)
	c.IndentedJSON(http.StatusCreated, covers)
}

// getCoverByID locates the cover whose ID value matches the id
// parameter sent by the client, then returns that cover as a response.
func getCoverByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of covers, looking for
	// an album whose ID value matches the parameter.
	for _, a := range covers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cover not found"})
}
