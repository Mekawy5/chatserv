package chat

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatService struct {
	Repository *ChatRepository
}

func NewChatService(r *ChatRepository) *ChatService {
	return &ChatService{
		Repository: r,
	}
}

func (s *ChatService) Create() {

}
