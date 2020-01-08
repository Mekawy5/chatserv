package main

import (
	"github.com/Mekawy5/chatserv/conf"
	"github.com/Mekawy5/chatserv/registry"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()
	defer db.Close()

	rmqc := tools.NewRabbitClient()
	rmqc.SetUp()
	defer rmqc.Conn.Close()

	appApi := registry.InitApplicationApi(db)
	chatApi := registry.InitChatApi(db)
	msgController := registry.InitMessageController(db, rmqc)

	r := gin.Default()

	r.GET("/applications", appApi.GetAll)
	r.GET("/application/:id", appApi.Get)
	r.POST("/application", appApi.Create)

	r.GET("/chats", chatApi.GetAll)
	r.GET("/chat/:number", chatApi.Get)
	r.POST("/chat/:app_token/create", chatApi.Create)

	r.POST("/app/:token/chat/:number/create-message", msgController.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
