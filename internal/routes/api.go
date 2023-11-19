package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	public := router.Group("")
	{
		public.POST("users/register", handlers.AuthHandler.Register)
		public.POST("users/login", handlers.AuthHandler.Login)
	}

	user := router.Group("").Use(middleware.IsValidJWT())
	{
		user.PATCH("users/topup", middleware.SetUserID(), handlers.AuthHandler.Topup)
	}

	return router
}
