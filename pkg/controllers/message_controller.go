package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mekawy5/chatserv/pkg/message"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	Service  *message.MessageService
	RabbitMQ *tools.RabbitClient
}

func NewMessageController(a *message.MessageService, rmqc *tools.RabbitClient) *MessageController {
	return &MessageController{
		Service:  a,
		RabbitMQ: rmqc,
	}
}

// func (ca *ChatApi) GetAll(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"chat": "Get All Chats"})
// }

// func (ca *ChatApi) Get(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"chat": "Get One Chats"})
// }

func (mc *MessageController) Create(c *gin.Context) {
	var msg message.Message
	err := c.BindJSON(&msg)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	chatNum, _ := strconv.Atoi(c.Param("number"))

	newMsg := mc.Service.Create(message.NewMessage(msg), c.Param("token"), uint(chatNum))

	mc.RabbitMQ.Publish(tools.MSGEXC, tools.MSGKEY, util.ToJson(newMsg))

	c.JSON(http.StatusOK, gin.H{"msg": message.GetMessage(newMsg)})
}
