package models

type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUserPwd struct {
	OldPassword string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type AuthPwd struct {
	Application string `json:"application" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Key         string `json:"key" binding:"required"`
}

type AuthGetPwd struct {
	Application string `json:"application" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Key         string `json:"key" binding:"required"`
}
