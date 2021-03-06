package application

import (
	"time"

	"github.com/Mekawy5/chatserv/pkg/chat"
)

type Application struct {
	ID         uint        `json:"id,omitempty"`
	Name       string      `json:"name"`
	Token      string      `json:"token"`
	ChatsCount uint        `json:"chats_count"`
	Chats      []chat.Chat `json:"chats"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func GetApplication(app ApplicationModel) Application {
	return Application{
		Name:       app.Name,
		Token:      app.Token,
		ChatsCount: app.ChatsCount,
		Chats:      chat.GetChats(app.Chats),
		CreatedAt:  app.CreatedAt,
		UpdatedAt:  app.UpdatedAt,
	}
}

func GetApplications(apps []ApplicationModel) []Application {
	var appDtos []Application
	for _, app := range apps {
		appDtos = append(appDtos, GetApplication(app))
	}
	return appDtos
}

func NewApplication(app Application) ApplicationModel {
	return ApplicationModel{
		Name:  app.Name,
		Token: app.Token,
	}
}
