package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// post user value
	// curl localhost:8080/user -XPOST -d '{"name":"James", "value": "Fox"}' -H 'Content-Type:application/json'
	r.POST("/user", func(c *gin.Context) {

		// Parse JSON
		var json struct {
			Name  string `json:"name" binding:"required"`
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[json.Name] = json.Value
			c.JSON(http.StatusOK, json)
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "3000"
	}
	r.Run("0.0.0.0:" + port)
}
