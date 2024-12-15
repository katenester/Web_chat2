package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/katenester/Web_chat2/backend/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.sigUp)
		auth.POST("/login", h.sigIp)
	}
	router.GET("/chats", h.getAllChats)
	router.POST("/chat/:user_name", h.createChat)
	router.POST("/chats/messages/:user_name", h.sendMessage)
	router.GET("/chats/messages/:user_name", h.getAllMessage)
	return router
}
