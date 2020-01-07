package chat

import (
	"time"
)

type Chat struct {
	ID            uint      `json:"id,omitempty"`
	ApplicationID uint      `json:"application_id,omitempty"`
	Number        uint      `json:"number"`
	MessagesCount uint      `json:"messages_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetChat(chat ChatModel) Chat {
	return Chat{
		Number:        chat.Number,
		MessagesCount: chat.MessagesCount,
		CreatedAt:     chat.CreatedAt,
		UpdatedAt:     chat.UpdatedAt,
	}
}

func GetChats(chats []ChatModel) []Chat {
	var chatDtos []Chat
	for _, chat := range chats {
		chatDtos = append(chatDtos, GetChat(chat))
	}
	return chatDtos
}

func NewChat(chat Chat) ChatModel {
	return ChatModel{
		Number:        chat.Number,
		ApplicationID: chat.ApplicationID,
	}
}
