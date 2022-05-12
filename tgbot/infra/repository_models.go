package infra

import (
	"time"

	"gorm.io/gorm"
)

type (
	Message struct {
		gorm.Model
		ChatID      int64 `gorm:"uniqueIndex:idx_msg"`
		MessageId   int   `gorm:"uniqueIndex:idx_msg"`
		MessageTime time.Time
		Message     string
	}
	Chat struct {
		gorm.Model
		// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
		Type string
		// Title for supergroups, channels and group chats
		Title string
		// UserName for private chats, supergroups and channels if available
		UserName string
	}
)
