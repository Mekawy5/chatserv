package chat

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{
		DB: db,
	}
}

func (a *ChatRepository) Save(chat ChatModel) ChatModel {
	a.DB.Save(&chat)
	return chat
}

func (a *ChatRepository) GetAppId(token string) uint {
	var app struct{ ID uint }
	a.DB.Table("applications").Select("id").Where("token = ?", token).Scan(&app)
	return app.ID
}
