package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardikkheni/seedsuite-server/requests"
)

func Login(c *gin.Context) {
	data, _ := c.Get("data")
	c.JSON(http.StatusOK, data.(*requests.LoginRequest))
}
