package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/katenester/Web_chat2/backend/internal/service"
	"path/filepath"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	publicPath := filepath.Join(".", "frontend")
	router.StaticFile("/", filepath.Join(publicPath, "index.html"))
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.sigIp)
	}
	api := router.Group("/chat", h.userIdentity)
	{
		api.GET("/", h.getAllChats)
		api.POST("/:user_name", h.createChat)
		api.GET("/messages/:user_name", h.getAllMessage)
		api.POST("/messages/:user_name", h.sendMessage)
	}
	return router
}
