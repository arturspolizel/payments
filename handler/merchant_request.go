package handler

import "github.com/arturspolizel/payments/model"

type MerchantCreateRequest struct {
	Name                    string `json:"name" binding:"required"`
	Active                  bool   `json:"active" binding:"required"`
	MaximumTransactionValue *uint  `json:"maximumTransactionValue" binding:"required"`
}

func (mcr *MerchantCreateRequest) toMerchant() model.Merchant {
	merchant := model.Merchant{}
	merchant.Name = mcr.Name
	merchant.Active = mcr.Active
	merchant.MaximumTransactionValue = mcr.MaximumTransactionValue

	return merchant
}
