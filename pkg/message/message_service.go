package message

import (
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MessageService struct {
	Repository *MessageRepository
}

func NewMessageService(r *MessageRepository) *MessageService {
	return &MessageService{
		Repository: r,
	}
}

func (s *MessageService) Create(msg MessageModel, appToken string, chatNum uint) MessageModel {
	appChatKey := util.GenerateAppChatKey(appToken, chatNum)
	redisClient := tools.NewRedisClient()
	lastMsgNum, chatId := redisClient.GetAppChatInfo(appChatKey)

	// TODO get last msg num & chat number from other service that will be connected to db.

	if lastMsgNum == 0 {
		msg.Number = 1
	} else {
		msg.Number = lastMsgNum + 1
	}

	if chatId == 0 {
		msg.ChatID = s.Repository.GetChatId(appToken, chatNum)
	} else {
		msg.ChatID = chatId
	}

	redisClient.SetAppChatInfo(appChatKey, msg.Number, msg.ChatID)

	return msg
}
