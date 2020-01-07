package chat

import (
	"log"
	"os"

	"github.com/Mekawy5/chatserv/tools"
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

func (s *ChatService) Create(chat ChatModel, appToken string) ChatModel {
	redisClient := tools.NewRedisClient()
	chatNum := redisClient.GetAppChatNumber(appToken)
	chat.Number = uint(chatNum + 1)

	appId := s.Repository.GetAppId(appToken)
	chat.ApplicationID = appId

	newChat := s.Repository.Save(chat)

	redisClient.SetAppChatNumber(appToken, chat.Number)
	return newChat
}

func handleErrors(errs []error) {
	if errs != nil {
		for err := range errs {
			log.Panicln(err)
		}
		os.Exit(1)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
