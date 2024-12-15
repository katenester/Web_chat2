package transport

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) sigUp(c *gin.Context) {

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) sigIp(c *gin.Context) {

}
