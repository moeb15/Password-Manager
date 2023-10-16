package models

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthPwd struct {
	Application string `json:"application" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Key         string `json:"key" binding:"required"`
}
