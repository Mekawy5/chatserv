package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChatApi struct {
	Service *ChatService
}

func NewChatApi(a *ChatService) *ChatApi {
	return &ChatApi{
		Service: a,
	}
}

func (ca *ChatApi) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"chat": "Get All Chats"})
}

func (ca *ChatApi) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"chat": "Get One Chats"})
}

func (ca *ChatApi) Create(c *gin.Context) {
	var chat Chat
	err := c.BindJSON(&chat)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}
