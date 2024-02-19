package handler

import (
	"github.com/arturspolizel/payments/pkg/auth/interfaces"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	paymentController interfaces.UserController
	router            *gin.RouterGroup
}

func NewUserHandler(paymentController interfaces.UserController, router *gin.RouterGroup) *UserHandler {
	return &UserHandler{
		paymentController: paymentController,
		router:            router,
	}
}

func (h *UserHandler) SetRouters() {
	h.router.POST("/login", h.Login)
	h.router.POST("/register", h.Register)
	h.router.POST("/validate", h.Validate)
}

func (h *UserHandler) Login(c *gin.Context) {
	panic("Not implemented")
}

func (h *UserHandler) Register(c *gin.Context) {
	panic("Not implemented")
}

func (h *UserHandler) Validate(c *gin.Context) {
	panic("Not implemented")
}
