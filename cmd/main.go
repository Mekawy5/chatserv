package main

import (
	"github.com/Mekawy5/chatserv/conf"
	"github.com/Mekawy5/chatserv/registry"
	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()
	defer db.Close()

	appApi := registry.InitApplicationApi(db)
	chatApi := registry.InitChatApi(db)

	r := gin.Default()

	r.GET("/applications", appApi.GetAll)
	r.GET("/application/:id", appApi.Get)
	r.POST("/application", appApi.Create)

	r.GET("/chats", chatApi.GetAll)
	r.GET("/chat/:number", chatApi.Get)
	r.POST("/chat/:app_token/create", chatApi.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
