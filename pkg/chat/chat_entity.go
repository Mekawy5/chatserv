package chat

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatModel struct {
	gorm.Model
	ApplicationID uint
	Number        uint
	MessagesCount uint
}

func (ChatModel) TableName() string {
	return "chats"
}
