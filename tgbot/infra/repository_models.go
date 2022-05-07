package infra

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func CreateMessageFromBot(message *tgbotapi.Message) *Message {
	return &Message{
		ChatID:      message.Chat.ID,
		MessageId:   message.MessageID,
		MessageTime: time.Unix(int64(message.Date), 0),
	}
}
