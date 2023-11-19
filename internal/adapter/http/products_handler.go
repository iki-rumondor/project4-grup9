package customHTTP

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ProductsHandler struct {
	Service *application.ProductsService
}

func NewProductsHandler(service *application.ProductsService) *ProductsHandler {
	return &ProductsHandler{
		Service: service,
	}
}

func (h *ProductsHandler) CreateProducts(c *gin.Context) {
	var body request.Products
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	product := domain.Products{
		Title:      body.Title,
		Price:      body.Price,
		Stock:      body.Stock,
		CategoryId: body.CategoryId,
	}

	result, err := h.Service.CreateProducts(&product)
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

	response := response.Products{
		ID:           result.ID,
		Title:        result.Title,
		Price:        result.Price,
		Stock:        result.Stock,
		CategoriesId: result.CategoryId,
		CreatedAt:    result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *ProductsHandler) GetProducts(c *gin.Context) {
	productID := c.GetUint("productID")
	result, err := h.Service.GetProducts(productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var products = []*response.Products{}

	for _, product := range *result {
		products = append(products, &response.Products{
			ID:           product.ID,
			Title:        product.Title,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoriesId: product.CategoryId,
			CreatedAt:    product.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductsHandler) UpdateProducts(c *gin.Context) {
	var body request.Products
	if err := c.ShouldBindJSON(&body); err != nil {
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

	urlParam := c.Param("id")
	productsID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	products := domain.Products{
		ID:         uint(productsID),
		Title:      body.Title,
		Price:      body.Price,
		Stock:      body.Stock,
		CategoryId: body.CategoryId,
	}

	result, err := h.Service.UpdateProducts(&products)
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

	c.JSON(http.StatusOK, response.UpdateProducts{
		ID:           result.ID,
		Title:        result.Title,
		Price:        result.Price,
		Stock:        result.Stock,
		CategoriesId: result.CategoryId,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	})
}

func (h *ProductsHandler) DeleteProducts(c *gin.Context) {
	urlParam := c.Param("id")
	productsID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	product := domain.Products{
		ID: uint(productsID),
	}

	if err := h.Service.DeleteProducts(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Category has been successfully deleted",
	})
}
