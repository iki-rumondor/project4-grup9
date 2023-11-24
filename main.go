package main

import (
	"log"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// autoMigration(gormDB)

	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	products_repo := repository.NewProductsRepository(gormDB)
	products_service := application.NewProductsService(products_repo)
	products_handler := customHTTP.NewProductsHandler(products_service)

	categories_repo := repository.NewCategoriesRepository(gormDB)
	categories_service := application.NewCategoriesService(categories_repo)
	categories_handler := customHTTP.NewCategoriesHandler(categories_service)

	transaction_repo := repository.NewTransactionRepository(gormDB)
	transaction_service := application.NewTransactionService(transaction_repo)
	transaction_handler := customHTTP.NewTransactionHandler(transaction_service)

	handlers := &customHTTP.Handlers{
		AuthHandler:        auth_handler,
		CategoriesHandler:  categories_handler,
		ProductsHandler:    products_handler,
		TransactionHandler: transaction_handler,
	}

	utils.NewCustomValidator(gormDB)
	var PORT = envPortOr("3000")
	routes.StartServer(handlers).Run(PORT)
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func autoMigration(db *gorm.DB) {

	db.Migrator().DropTable(&domain.User{})
	db.Migrator().CreateTable(&domain.User{})

	db.Migrator().DropTable(&domain.Categories{})
	db.Migrator().CreateTable(&domain.Categories{})

	db.Migrator().DropTable(&domain.Products{})
	db.Migrator().CreateTable(&domain.Products{})

	db.Migrator().DropTable(&domain.TransactionHistory{})
	db.Migrator().CreateTable(&domain.TransactionHistory{})

	db.Create(&domain.User{
		FullName: "Administrator",
		Email: "admin@admin.com",
		Password: "123456",
		Role: "admin",
		Balance: 0,
	})
}
