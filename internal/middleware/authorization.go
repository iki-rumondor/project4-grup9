package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role := mapClaims["role"].(string)
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
				Success: false,
				Message: "access denied due to invalid credentials",
			})
			return
		}
		c.Next()
	}
}

func SetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		userID := uint(mapClaims["id"].(float64))

		c.Set("user_id", userID)
		c.Next()

	}
}
