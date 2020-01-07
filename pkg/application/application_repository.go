package application

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ApplicationRepository struct {
	DB *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{DB: db}
}

func (a *ApplicationRepository) GetAll() []ApplicationModel {
	var apps []ApplicationModel
	a.DB.Preload("Chats").Find(&apps)
	return apps
}

func (a *ApplicationRepository) Get(id uint) ApplicationModel {
	var app ApplicationModel
	a.DB.Find(&app, id)
	return app
}

func (a *ApplicationRepository) Save(app ApplicationModel) ApplicationModel {
	a.DB.Save(&app)
	return app
}

func (a *ApplicationRepository) GetAppIdByToken(token string) uint {
	var app ApplicationModel
	a.DB.Where("token = ?", token).First(&app)
	return app.ID
}
