package controller

import (
	"github.com/arturspolizel/payments/interfaces"
	"github.com/arturspolizel/payments/model"
)

type MerchantController struct {
	merchantRepository interfaces.MerchantRepository
}

func NewMerchantController(merchantRepository interfaces.MerchantRepository) *MerchantController {
	return &MerchantController{
		merchantRepository: merchantRepository,
	}
}

func (c *MerchantController) Get(id uint) model.Merchant {
	merchant := c.merchantRepository.Get(id)
	return merchant
}

func (c *MerchantController) Create(merchant model.Merchant) uint {
	id := c.merchantRepository.Create(merchant)
	return id
}
