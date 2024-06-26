package models

import (
	"time"
)

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Language  string    `json:"language" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	RoleID    uint      `json:"roleID" gorm:"not null"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
}

type UserFeedback struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"userID" gorm:"not null"`
	Comments    string    `json:"comments" gorm:"not null"`
	BotFeedback int       `json:"botFeedback" gorm:"not null;default:0;check:bot_feedback >= 1 AND bot_feedback <= 5"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	User        User      `user:"user" gorm:"foreignKey:UserID"`
}
