package message

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (a *MessageRepository) Save(msg MessageModel) MessageModel {
	a.DB.Save(&msg)
	return msg
}

func (a *MessageRepository) GetChatId(appToken string, chatNum uint) uint {
	var chat struct{ ID uint }
	a.DB.Table("chats").Select("chats.id").Joins("left join applications on chats.application_id = applications.id").Where("applications.token = ?", appToken).Where("chats.number = ?", chatNum).Scan(&chat)
	log.Println("DB hit to get chat id from chat number and app token")
	return chat.ID
}
