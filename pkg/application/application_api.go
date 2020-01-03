package application

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApplicationApi struct {
	Service *ApplicationService
}

func NewApplicationApi(a *ApplicationService) *ApplicationApi {
	return &ApplicationApi{
		Service: a,
	}
}

func (a *ApplicationApi) GetAll(c *gin.Context) {
	apps := a.Service.GetAll()
	c.JSON(http.StatusOK, gin.H{"apps": GetApplications(apps)})
}

func (a *ApplicationApi) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	app := a.Service.Get(uint(id))
	c.JSON(http.StatusOK, gin.H{"app": GetApplication(app)})
}

func (a *ApplicationApi) Create(c *gin.Context) {
	var app Application
	err := c.BindJSON(&app)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	newApp := a.Service.Save(NewApplication(app))
	c.JSON(http.StatusOK, gin.H{"app": GetApplication(newApp)})
}
