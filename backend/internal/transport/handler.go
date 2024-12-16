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

	// Раздача статических файлов из директории "frontend"
	router.Static("/static", "./frontend2")

	// Рендеринг HTML страниц из директории "frontend"
	router.GET("/", func(c *gin.Context) {
		c.File("./frontend2/index.html") // Главная страница
	})
	router.GET("/register", func(c *gin.Context) {
		c.File("./frontend2/register.html") // Страница регистрации
	})
	router.GET("/login", func(c *gin.Context) {
		c.File("./frontend2/login.html") // Страница авторизации
	})
	router.GET("/chats", func(c *gin.Context) {
		c.File("./frontend2/chats.html") // Страница чатов
	})

	// Группа маршрутов для аутентификации
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.sigIp)
	}

	// Группа маршрутов для чатов с авторизацией
	api := router.Group("/chat", h.userIdentity)
	{
		api.GET("/", h.getAllChats)
		api.POST("/:user_name", h.createChat)
		api.GET("/messages/:user_name", h.getAllMessage)
		api.POST("/messages/:user_name", h.sendMessage)
	}

	return router
}
