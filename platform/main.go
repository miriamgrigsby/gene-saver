package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type genetic struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	router := gin.Default()

	// Add a home endpoint for the genetics overview homepage
	router.GET("/", getGeneticsOverview)

	// todo add readiness and liveness probes endpoints

	router.GET("/health", health)
	router.GET("/readiness", readiness)

	router.Run("0.0.0.0:4040")
}

// overview responds with an overview of genetics, genes, and genetic therapy data
// TODO get genetics data, store it in db, and write adapter/handler to return that data here
var genetics = []genetic{
	{ID: "1", Title: "Blue Train"},
	{ID: "2", Title: "Jeru"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown"},
}

// getAlbums responds with the list of all albums as JSON.
func getGeneticsOverview(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, genetics)
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "IM ALIVE!"})
}
func readiness(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "IM READY!"})
}
