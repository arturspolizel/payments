package handler

import (
	"net/http"

	"github.com/arturspolizel/payments/interfaces"
	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchantController interfaces.MerchantController
	router             *gin.RouterGroup
}

func NewMerchantHandler(merchantController interfaces.MerchantController, router *gin.RouterGroup) *MerchantHandler {
	return &MerchantHandler{
		merchantController: merchantController,
		router:             router,
	}
}

func (h *MerchantHandler) SetRouters() {
	h.router.GET("/:id", h.GetMerchant)
	h.router.POST("/", h.CreateMerchant)
}

func (h *MerchantHandler) GetMerchant(c *gin.Context) {
	h.merchantController.Get(c.GetUint("id"))
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	merchant := MerchantCreateRequest{}

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !merchant.Currency.Validate() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency must be a valid ISO 4217 3-letter code"})
		return
	}

	id := h.merchantController.Create(merchant.toMerchant())
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}
