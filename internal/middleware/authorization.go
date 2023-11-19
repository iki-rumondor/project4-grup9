package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetString("jwt")
		if jwt == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Success: false,
				Message: "JWT token is not valid",
			})
			return
		}

		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		role := mapClaims["role"].(float64)
		if role != 1 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
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
