package message

import (
	"time"
)

type Message struct {
	ID        uint      `json:"id,omitempty"`
	ChatID    uint      `json:"chat_id,omitempty"`
	Number    uint      `json:"number"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetMessage(msg MessageModel) Message {
	return Message{
		Number:    msg.Number,
		Body:      msg.Body,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
	}
}

func GetMessages(msgs []MessageModel) []Message {
	var msgsDtos []Message
	for _, msg := range msgs {
		msgsDtos = append(msgsDtos, GetMessage(msg))
	}
	return msgsDtos
}

func NewMessage(msg Message) MessageModel {
	return MessageModel{
		ChatID: msg.ChatID,
		Number: msg.Number,
		Body:   msg.Body,
	}
}
