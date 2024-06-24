package v1

import "github.com/chenchi1009/go-starter-kit/internal/model"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

type User struct {
	ID       uint         `json:"id"`
	Username string       `json:"username" binding:"required"`
	Rules    []model.Rule `json:"rules"`
}

type UpdateProfileRequest struct {
	User
}

type GetProfileResponse struct {
	User
}

type ListUserRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type ListUserResponse struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}
