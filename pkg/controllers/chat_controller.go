package controllers

import (
	"net/http"
	"time"

	"github.com/Mekawy5/chatserv/pkg/chat"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	"github.com/gin-gonic/gin"
)

type ChatConroller struct {
	Q *tools.RabbitClient
}

func NewChatConroller(q *tools.RabbitClient) *ChatConroller {
	return &ChatConroller{
		Q: q,
	}
}

func (cc *ChatConroller) Create(c *gin.Context) {
	var ch chat.Chat
	err := c.BindJSON(&ch)
	util.HandleErrors(err, c)

	// append app token to struct, will be used in another service to fetch app id
	// append created at, updated at to track request date not database storing date
	ch.AppToken = c.Param("app_token")
	ch.CreatedAt = time.Now()
	ch.UpdatedAt = ch.CreatedAt

	cc.Q.Publish(tools.CHTEXC, tools.CHTKEY, util.ToJson(ch))
	c.JSON(http.StatusOK, gin.H{"chat": ch})
}
