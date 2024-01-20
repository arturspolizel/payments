package handler

import (
	"github.com/arturspolizel/payments/interfaces"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentController interfaces.PaymentController
	router            *gin.Engine
}

func NewPaymentHandler(paymentController interfaces.PaymentController, router *gin.Engine) *PaymentHandler {
	return &PaymentHandler{
		paymentController: paymentController,
		router:            router,
	}
}

func (h *PaymentHandler) SetRouters() {
	h.router.GET("/:id", h.GetPayment)

	h.router.POST("/", h.CreatePayment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	h.paymentController.Get(c.GetUint("id"))
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	//TODO: parse json, create object
}
