//+build wireinject

package registry

import (
	"github.com/Mekawy5/chatserv/pkg/application"
	"github.com/Mekawy5/chatserv/pkg/chat"
	"github.com/Mekawy5/chatserv/pkg/controllers"
	"github.com/Mekawy5/chatserv/pkg/message"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitApplicationApi(db *gorm.DB) *application.ApplicationApi {
	wire.Build(application.NewApplicationApi, application.NewApplicationService, application.NewApplicationRepository)
	return &application.ApplicationApi{}
}

func InitChatApi(db *gorm.DB) *chat.ChatApi {
	wire.Build(chat.NewChatApi, chat.NewChatService, chat.NewChatRepository)
	return &chat.ChatApi{}
}

func InitMessageController(db *gorm.DB, rmc *tools.RabbitClient) *controllers.MessageController {
	wire.Build(controllers.NewMessageController, message.NewMessageService, message.NewMessageRepository)
	return &controllers.MessageController{}
}

func InitChatController(*tools.RabbitClient) *controllers.ChatConroller {
	wire.Build(controllers.NewChatConroller)
	return &controllers.ChatConroller{}
}
