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

func (c *MerchantController) Get(id uint) (model.Merchant, error) {
	merchant, err := c.merchantRepository.Get(id)
	return merchant, err
}

func (c *MerchantController) Create(merchant model.Merchant) (uint, error) {
	id, err := c.merchantRepository.Create(merchant)
	return id, err
}
