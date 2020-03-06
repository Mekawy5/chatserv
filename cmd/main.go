package main

import (
	"github.com/Mekawy5/chatserv/registry"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	//TODO make this app stateless (remove db and redis, all persistent info needed will be consumed from the chatapp application)
	// db := conf.InitDB()
	// defer db.Close()

	rmqc := tools.NewRabbitClient()
	go rmqc.Setup()

	defer rmqc.Conn.Close()

	// appApi := registry.InitApplicationApi(db)
	// chatApi := registry.InitChatApi(db)

	appCtr := registry.InitApplicationController(rmqc)
	chatCtr := registry.InitChatController(rmqc)
	msgCtr := registry.InitMessageController(rmqc)

	r := gin.Default()

	// r.GET("/applications", appApi.GetAll)
	// r.GET("/application/:id", appApi.Get)
	// r.POST("/application", appApi.Create)
	r.POST("/application", appCtr.Create)

	// r.GET("/chats", chatApi.GetAll)
	// r.GET("/chat/:number", chatApi.Get)
	// r.POST("/chat/:app_token/create", chatApi.Create)
	r.POST("/chat/:app_token/create", chatCtr.Create)

	r.POST("/app/:token/chat/:number/create-message", msgCtr.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
