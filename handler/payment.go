package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arturspolizel/payments/interfaces"
	"github.com/arturspolizel/payments/model"
	"github.com/arturspolizel/payments/utils"
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
	h.router.GET("/", h.ListPayments)
	h.router.POST("/", h.CreatePayment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id, err := utils.PathUint(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Couldn't parse id parameter: %s", err.Error())})
		return
	}

	payment, err := h.paymentController.Get(id)

	if err != nil {
		var notFoundErr *model.ErrDatabaseNotFound
		if errors.As(err, &notFoundErr) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) ListPayments(c *gin.Context) {
	var query PaymentListRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid query parameters: %s", err.Error())})
		return
	}

	payments, err := h.paymentController.List(query.StartId, query.PageSize, query.StartDate, query.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var paginatedReturn PaginationResponse[model.Payment]
	if len(payments) > 0 {
		paginatedReturn = PaginationResponse[model.Payment]{
			StartId: payments[0].ID,
			EndId:   payments[len(payments)-1].ID,
			Count:   uint(len(payments)),
			Data:    payments,
		}
	} else {
		paginatedReturn = PaginationResponse[model.Payment]{
			Data: payments,
		}
	}

	c.JSON(http.StatusOK, paginatedReturn)
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

	id, err := h.paymentController.Create(payment.toPayment())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
