package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "anhkhoi")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong ... ping" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "anhkhoi"},
	})
}
