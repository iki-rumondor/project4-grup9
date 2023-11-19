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

		user.GET("/products", handlers.ProductsHandler.GetProducts)

		user.POST("/transactions", handlers.TransactionHandler.CreateTransaction)
		user.GET("/transactions/my-transactions", handlers.TransactionHandler.GetMyTransaction)
	}

	admin := router.Group("").Use(middleware.IsValidJWT(), middleware.IsAdmin())
	{
		admin.POST("/categories", handlers.CategoriesHandler.CreateCategories)
		admin.GET("/categories", handlers.CategoriesHandler.GetCategories)
		admin.PATCH("/categories/:id", handlers.CategoriesHandler.UpdateCategories)
		admin.DELETE("/categories/:id", handlers.CategoriesHandler.DeleteCategories)

		admin.POST("/products", handlers.ProductsHandler.CreateProducts)
		admin.PUT("/products/:id", handlers.ProductsHandler.UpdateProducts)
		admin.DELETE("/products/:id", handlers.ProductsHandler.DeleteProducts)

		admin.GET("/transactions/user-transactions", handlers.TransactionHandler.GetUserTransaction)
	}

	return router
}
