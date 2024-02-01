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
	h.router.GET("/merchant/:id", h.GetMerchant)
	h.router.POST("/merchant/", h.CreateMerchant)
}

func (h *MerchantHandler) GetMerchant(c *gin.Context) {
	id, err := utils.PathUint(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Couldn't parse id parameter: %s", err.Error())})
		return
	}

	merchant, err := h.merchantController.Get(uint(id))
	if err != nil {
		var notFoundErr *model.ErrDatabaseNotFound
		if errors.As(err, &notFoundErr) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}
	c.JSON(http.StatusAccepted, merchant)
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	merchant := MerchantCreateRequest{}

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.merchantController.Create(merchant.toMerchant())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}