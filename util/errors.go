package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrors(err error, c *gin.Context) {
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
}
