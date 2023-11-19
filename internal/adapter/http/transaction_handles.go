package customHTTP

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	Service *application.TransactionService
}

func NewTransactionHandler(service *application.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		Service: service,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var body request.Transaction
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "your request body is not valid",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	transaction := domain.TransactionHistory{
		ProductsId: body.ProductsId,
		Quantity:   body.Quantity,
	}

	result, err := h.Service.CreateTransaction(&transaction)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.TransactionBill{
		Total_Price:    result.Total_Price,
		Quantity:       result.Quantity,
		Products_Title: result.Products.Title,
	}

	successMessage := "You have successfully purchased the product"

	c.JSON(http.StatusCreated, gin.H{
		"message":          successMessage,
		"transaction_bill": response,
	})
}

func (h *TransactionHandler) GetMyTransaction(c *gin.Context) {
	result, err := h.Service.GetMyTransaction()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var transaction = []*response.MyTransaction{}

	for _, transactions := range *result {
		transaction = append(transaction, &response.MyTransaction{
			ID:          transactions.ID,
			ProductsId:  transactions.ProductsId,
			UserId:      transactions.UserId,
			Quantity:    transactions.Quantity,
			Total_Price: transactions.Total_Price,
			Product: response.Product{
				ID:           transactions.Products.ID,
				Title:        transactions.Products.Title,
				Price:        transactions.Products.Price,
				Stock:        transactions.Products.Stock,
				CategoriesId: transactions.Products.CategoryId,
				CreatedAt:    transactions.Products.CreatedAt,
				UpdatedAt:    transactions.Products.UpdatedAt,
			},
		})
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetUserTransaction(c *gin.Context) {
	result, err := h.Service.GetUserTransaction()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var transaction = []*response.UserTransaction{}

	for _, transactions := range *result {
		transaction = append(transaction, &response.UserTransaction{
			ID:          transactions.ID,
			ProductsId:  transactions.ProductsId,
			UserId:      transactions.UserId,
			Quantity:    transactions.Quantity,
			Total_Price: transactions.Total_Price,
			Product: response.Product{
				ID:           transactions.Products.ID,
				Title:        transactions.Products.Title,
				Price:        transactions.Products.Price,
				Stock:        transactions.Products.Stock,
				CategoriesId: transactions.Products.CategoryId,
				CreatedAt:    transactions.Products.CreatedAt,
				UpdatedAt:    transactions.Products.UpdatedAt,
			},
			Users: response.Users{
				ID:        transactions.User.ID,
				Email:     transactions.User.Email,
				FullName:  transactions.User.FullName,
				Balance:   transactions.User.Balance,
				CreatedAt: transactions.User.CreatedAt,
				UpdatedAt: transactions.User.UpdatedAt,
			},
		})
	}

	c.JSON(http.StatusOK, transaction)
}
