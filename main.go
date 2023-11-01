package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Position struct {
	Xval int `json:"xval"`
	Yval int `json:"yval"`
}

// getPosition responds with JSON.
func getPosition(c *gin.Context) {
	var position Position = Position{Xval: rand.Intn(10), Yval: rand.Intn(10)}
	c.IndentedJSON(http.StatusOK, position)
}

func main() {
	router := gin.Default()
	router.GET("/position", getPosition)

	router.Run("localhost:8080")
}
