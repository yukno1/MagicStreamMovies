package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukno1/MagicStreamMovies/Server/MagicStreamMoviesServer/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	// endpoint called before any publicly exposed endppints are called
	return func(c *gin.Context) {
		token, err := utils.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort() // no continue, abort further execution
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)

		// continue to execute the targeted endpoint
		c.Next()
	}
}
