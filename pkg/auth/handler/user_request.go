package handler

import "github.com/arturspolizel/payments/pkg/auth/model"

type UserCreateRequest struct {
	Email      string `json:"email" binding:"required"`
	Name       string `json:"name" binding:"required"`
	MerchantId uint   `json:"merchantId" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserInfo struct {
	Email      string           `json:"email" binding:"required"`
	Name       string           `json:"name" binding:"required"`
	MerchantId uint             `json:"merchantId" binding:"required"`
	Status     model.UserStatus `json:"status"`
}

func (ucr *UserCreateRequest) toUser() model.User {
	user := model.User{}
	user.Name = ucr.Name
	user.Email = ucr.Email
	user.MerchantId = ucr.MerchantId

	return user
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ValidateRequest struct {
	Code string `form:"code" binding:"required"`
}
