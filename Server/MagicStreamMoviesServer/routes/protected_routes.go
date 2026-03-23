package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yukno1/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/yukno1/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controllers.GetMovie())
	router.POST("/addmovie", controllers.AddMovie())
}
