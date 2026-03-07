package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yukno1/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	router.GET("/movies", controllers.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
