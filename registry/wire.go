//+build wireinject

package registry

import (
	"github.com/Mekawy5/chatserv/pkg/controllers"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/google/wire"
)

// func InitApplicationApi(db *gorm.DB) *application.ApplicationApi {
// 	wire.Build(application.NewApplicationApi, application.NewApplicationService, application.NewApplicationRepository)
// 	return &application.ApplicationApi{}
// }

// func InitChatApi(db *gorm.DB) *chat.ChatApi {
// 	wire.Build(chat.NewChatApi, chat.NewChatService, chat.NewChatRepository)
// 	return &chat.ChatApi{}
// }

func InitApplicationController(rmc *tools.RabbitClient) *controllers.ApplicationController {
	wire.Build(controllers.NewApplicationController)
	return &controllers.ApplicationController{}
}

func InitMessageController(rmc *tools.RabbitClient) *controllers.MessageController {
	wire.Build(controllers.NewMessageController)
	return &controllers.MessageController{}
}

func InitChatController(*tools.RabbitClient) *controllers.ChatConroller {
	wire.Build(controllers.NewChatConroller)
	return &controllers.ChatConroller{}
}
