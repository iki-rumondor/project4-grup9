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

type CategoriesHandler struct {
	Service *application.CategoriesService
}

func NewCategoriesHandler(service *application.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		Service: service,
	}
}

func (h *CategoriesHandler) CreateCategories(c *gin.Context) {

	var body request.Categories
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

	categories := domain.Categories{
		Type: body.Type,
	}

	result, err := h.Service.CreateCategories(&categories)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreateCategories{
		ID:                   result.ID,
		Type:                 result.Type,
		Sold_Product_Ammount: result.Sold_Product_Ammount,
		Created_At:           result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CategoriesHandler) GetCategories(c *gin.Context) {

	result, err := h.Service.GetCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var categories = []*response.Categories{}

	for _, category := range *result {

		var Products = []*response.ProductsCategories{}

		for _, product := range category.Products {
			Products = append(Products, &response.ProductsCategories{
				ID:        product.ID,
				Title:     product.Title,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			})
		}

		categories = append(categories, &response.Categories{
			ID:                   category.ID,
			Type:                 category.Type,
			Sold_Product_Ammount: category.Sold_Product_Ammount,
			CreatedAt:            category.CreatedAt,
			UpdatedAt:            category.UpdatedAt,
			ProductsCategories:   Products,
		})
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoriesHandler) UpdateCategories(c *gin.Context) {
	var body request.Categories
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

	urlParam := c.Param("id")
	categoriesID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	categories := domain.Categories{
		ID:   uint(categoriesID),
		Type: body.Type,
	}

	result, err := h.Service.UpdateCategories(&categories)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateCategories{
		ID:                   result.ID,
		Type:                 result.Type,
		Sold_Product_Ammount: result.Sold_Product_Ammount,
		UpdatedAt:            result.UpdatedAt,
	})
}

func (h *CategoriesHandler) DeleteCategories(c *gin.Context) {

	urlParam := c.Param("id")
	categoryID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	category := domain.Categories{
		ID: uint(categoryID),
	}

	if err := h.Service.DeleteCategories(&category); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Categories has been successfully deleted",
	})
}
