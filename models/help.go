package models

import (
	"time"
)

type HelpMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoleID    uint      `json:"roleID" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
}

type HelpMessageTranslation struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	HelpMessageID uint        `json:"helpID" gorm:"not null"`
	Language      string      `json:"language" gorm:"not null"`
	Content       string      `json:"content" gorm:"not null"`
	CreatedAt     time.Time   `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time   `gorm:"default:current_timestamp"`
	HelpMessage   HelpMessage `gorm:"foreignKey:HelpMessageID"`
}
