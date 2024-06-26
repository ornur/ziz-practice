package models

import (
	"time"
)

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoleID    uint      `json:"roleID" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
}

type MessageTranslation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MessageID uint      `json:"messageID" gorm:"not null"`
	Language  string    `json:"language" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Message   Message   `json:"message" gorm:"foreignKey:MessageID"`
}
