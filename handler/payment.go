package handler

import (
	"net/http"

	"github.com/arturspolizel/payments/interfaces"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentController interfaces.PaymentController
	router            *gin.RouterGroup
}

func NewPaymentHandler(paymentController interfaces.PaymentController, router *gin.RouterGroup) *PaymentHandler {
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
	payment := PaymentCreateRequest{}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !payment.Currency.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency must be a valid ISO 4217 3-letter code"})
		return
	}

	id := h.paymentController.Create(payment.toPayment())
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}
