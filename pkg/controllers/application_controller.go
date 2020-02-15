package controllers

import (
	"net/http"
	"time"

	"github.com/Mekawy5/chatserv/pkg/application"
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

type ApplicationController struct {
	Q *tools.RabbitClient
}

func NewApplicationController(q *tools.RabbitClient) *ApplicationController {
	return &ApplicationController{
		Q: q,
	}
}

func (ac *ApplicationController) Create(c *gin.Context) {
	var app application.Application
	err := c.BindJSON(&app)
	util.HandleErrors(err, c)

	// append created at, updated at to track request date not database storing date
	app.Token = shortuuid.New()
	app.CreatedAt = time.Now()
	app.UpdatedAt = app.CreatedAt

	ac.Q.Publish(tools.APPEXC, tools.APPKEY, util.ToJson(app))
	c.JSON(http.StatusOK, gin.H{"app": app})
}
