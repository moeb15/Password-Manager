package models

type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthPwd struct {
	Application string `json:"application" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Key         string `json:"key" binding:"required"`
}

type AuthGetPwd struct {
	Application string `json:"application" binding:"required"`
	Key         string `json:"key" binding:"required"`
}
