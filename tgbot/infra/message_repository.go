package infra

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	MessageRepository struct {
		db      *gorm.DB
		logger  *logrus.Entry
		chatIds map[uint]bool
		lock    sync.RWMutex
	}
)

func GetSQLiteDB(connectionString string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
}

func CreateMessageRepository(db *gorm.DB) *MessageRepository {
	repository := &MessageRepository{
		db:     db,
		logger: GetLogger(fmt.Sprintf("repository:%s", db.Dialector.Name())),
	}
	repository.logger.Info("init")

	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Chat{})
	return repository
}

func (mr *MessageRepository) Save(message *Message) error {
	return mr.db.Save(message).Error
}

func (mr *MessageRepository) GetChats() ([]*Chat, error) {
	mr.lock.RLock()
	defer mr.lock.RUnlock()

	var chats []*Chat
	db := mr.db.Find(&chats)
	chatIds := make(map[uint]bool)
	for _, chat := range chats {
		chatIds[chat.ID] = true
	}
	mr.chatIds = chatIds

	return chats, db.Error
}
func (mr *MessageRepository) updateChatIds() {
	rows, err := mr.db.Table("chats").Select("ID").Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id uint
		rows.Scan(&id)
		mr.chatIds[id] = true
	}
}
func (mr *MessageRepository) SaveChat(chat *tgbotapi.Chat) error {
	if len(mr.chatIds) == 0 {
		mr.updateChatIds()
	}
	if _, ok := mr.chatIds[uint(chat.ID)]; ok {
		return nil
	}
	mr.lock.Lock()
	defer mr.lock.Unlock()
	newChat := &Chat{
		Model:    gorm.Model{ID: uint(chat.ID)},
		Type:     chat.Type,
		Title:    chat.Title,
		UserName: chat.UserName,
	}
	mr.chatIds[newChat.ID] = true
	return mr.db.Save(newChat).Error
}

func (mr *MessageRepository) SaveMessage(message *tgbotapi.Message) error {
	mr.lock.Lock()
	defer mr.lock.Unlock()
	messageBody, _ := json.Marshal(message)
	newMessage := &Message{
		Model:       gorm.Model{ID: uint(message.MessageID)},
		ChatID:      message.Chat.ID,
		MessageId:   message.MessageID,
		MessageTime: time.Unix(int64(message.Date), 0),
		Message:     string(messageBody),
	}
	return mr.db.Save(newMessage).Error
}

func (mr *MessageRepository) HouseKeeping(maxAge time.Duration) error {
	mr.lock.Lock()
	defer mr.lock.Unlock()
	logger := GetLogger("repository:housekeeping")
	beforeDate := time.Now().Add(-maxAge)
	logger.Infof("housekeeping data before %s", beforeDate)
	res := mr.db.Unscoped().Where("updated_at < ?", beforeDate).Delete(&Message{})
	if err := res.Error; err != nil {
		logger.Errorf("failed to delete messages: %s", err)
		return err
	}

	logger.Infof("deleted %d messages", res.RowsAffected)
	res = mr.db.Unscoped().Where("updated_at < ?", beforeDate).Delete(&Chat{})
	if err := res.Error; err != nil {
		logger.Errorf("failed to delete chats: %s", err)
		return err
	}
	logger.Infof("deleted %d chats", res.RowsAffected)
	return nil
}
