package application

import (
	"github.com/lithammer/shortuuid"
	"time"
)

type Application struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetApplication(app ApplicationModel) Application {
	return Application{
		ID:        app.ID,
		Name:      app.Name,
		Token:     app.Token,
		CreatedAt: app.CreatedAt,
		UpdatedAt: app.UpdatedAt,
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
		Token: shortuuid.New(),
	}
}
