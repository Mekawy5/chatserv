package message

import "github.com/jinzhu/gorm"

type MessageModel struct {
	gorm.Model
	ChatID uint
	Body   string
	Number uint
}

func (MessageModel) TableName() string {
	return "messages"
}
