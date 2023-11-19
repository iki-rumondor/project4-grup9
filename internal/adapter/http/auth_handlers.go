package customHTTP

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var body request.Register
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user := domain.User{
		FullName: body.FullName,
		Email:    body.Email,
		Password: body.Password,
		Role:     "customer",
	}

	res, err := h.Service.CreateUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.CreatedUser{
		ID:       res.ID,
		FullName: res.FullName,
		Email:    res.Email,
		Password: res.Password,
		Balance:  res.Balance,
		CreatedAt: res.CreatedAt,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body request.Login
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user := domain.User{
		Email:    body.Email,
		Password: body.Password,
	}

	jwt, err := h.Service.VerifyUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func (h *AuthHandler) Topup(c *gin.Context) {
	var body request.Topup
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}


	user := domain.User{
		ID:    c.GetUint("user_id"),
		Balance: body.Balance,
	}

	result, err := h.Service.UpdateBalance(&user)
	

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Success: false,
				Message: err.Error(),
			})
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp.%d", result.Balance),
	})
}

