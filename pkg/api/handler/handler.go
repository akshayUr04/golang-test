package handler

import (
	"golang-test/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequest(c *gin.Context) {

	ch := make(chan map[string]interface{})
	go helpers.Worker(ch)
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ch <- request

}
