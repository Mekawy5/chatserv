package conf

import (
	"os"

	"github.com/Mekawy5/chatserv/pkg/application"
	"github.com/Mekawy5/chatserv/pkg/chat"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	var dbUrl string
	if url := os.Getenv("DB_URL"); url == "" {
		dbUrl = "root:123@tcp(database:3306)/chat?charset=utf8&parseTime=True&loc=Local"
	} else {
		dbUrl = url
	}

	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&application.ApplicationModel{})
	db.AutoMigrate(&chat.ChatModel{}).AddForeignKey("application_id", "applications(id)", "CASCADE", "CASCADE")

	return db
}
