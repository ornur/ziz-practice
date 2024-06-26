package models

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Language string `json:"language" binding:"required"`
	RoleID   uint   `json:"roleID" binding:"required"`
}