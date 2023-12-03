package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mpopadic/go-fly/db"
	"github.com/mpopadic/go-fly/handlers"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	_db = db.Init()
}

func setupRouter() *gin.Engine {
	dbHandler := handlers.New(_db)
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// // Get user value
	r.GET("/user/:id", dbHandler.GetUser)

	// post user value
	// curl localhost:8080/user -XPOST -d '{"name":"James", "value": "Fox"}' -H 'Content-Type:application/json'
	r.POST("/user", dbHandler.AddUser)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	r.Run("0.0.0.0:" + port)
}
