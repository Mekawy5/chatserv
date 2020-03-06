package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Mekawy5/chatserv/pkg/message"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	Q *tools.RabbitClient
}

func NewMessageController(q *tools.RabbitClient) *MessageController {
	return &MessageController{
		Q: q,
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

	// append app token, chat number to struct, will be used in another service to fetch & cache app id and the proper chat last msg number and the chat id ..etc
	// append created at, updated at to track request date not database storing date
	// newMsg := mc.Service.Create(message.NewMessage(msg), c.Param("token"), uint(chatNum))
	msg.AppToken = c.Param("token")
	msg.ChatNum = uint(chatNum)
	msg.CreatedAt = time.Now()
	msg.UpdatedAt = msg.CreatedAt

	mc.Q.Publish(tools.MSGEXC, tools.MSGKEY, util.ToJson(msg))

	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
