package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arturspolizel/payments/pkg/auth/interfaces"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/arturspolizel/payments/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userController interfaces.UserController
	router         *gin.RouterGroup
}

func NewUserHandler(paymentController interfaces.UserController, router *gin.RouterGroup) *UserHandler {
	return &UserHandler{
		userController: paymentController,
		router:         router,
	}
}

func (h *UserHandler) SetRouters() {
	h.router.POST("/login", h.Login)
	h.router.POST("/register", h.Register)
	h.router.POST("/validate", h.Validate)
}

func (h *UserHandler) Login(c *gin.Context) {
	login := LoginRequest{}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userController.Login(login.Email, login.Password)
	if err != nil {
		var notFoundErr *utils.ErrDatabaseNotFound
		var unauthorizedErr *model.ErrAuthenticationFailed
		var invalidStatusErr *model.ErrInvalidUserStatus
		if errors.As(err, &notFoundErr) || errors.As(err, &unauthorizedErr) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else if errors.As(err, &invalidStatusErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) Register(c *gin.Context) {
	userRequest := UserCreateRequest{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.userController.Create(userRequest.toUser(), userRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *UserHandler) Validate(c *gin.Context) {
	var query ValidateRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid query parameters: %s", err.Error())})
		return
	}

	err = h.userController.Validate(query.Code)
	if err != nil {
		var notFoundErr *utils.ErrDatabaseNotFound
		var invalidCode *model.ErrInvalidEmailCode
		var invalidStatusErr *model.ErrInvalidUserStatus
		if errors.As(err, &notFoundErr) || errors.As(err, &invalidCode) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if errors.As(err, &invalidStatusErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, nil)
}
