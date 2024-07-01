package models

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserInput struct {
	Telegram_ID uint   `json:"telegramID" binding:"required"`
	Language    string `json:"language" binding:"required"`
	Email       string `json:"email" binding:"required"`
	RoleID      uint   `json:"roleID" binding:"required"`
}

type UpdateUserInput struct {
	Telegram_ID uint   `json:"telegramID"`
	Language    string `json:"language"`
	Email       string `json:"email"`
	RoleID      uint   `json:"roleID"`
}

type UserFeedbackInput struct {
	UserID      uint   `json:"userID" binding:"required"`
	Comments    string `json:"comments"`
	BotFeedback int    `json:"botFeedback"`
}
