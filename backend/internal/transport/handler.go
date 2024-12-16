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

	// Раздача статических файлов из директории "frontend"
	//router.Static("/static", "./frontend2")
	//
	//// Рендеринг HTML страниц из директории "frontend"
	//router.GET("/", func(c *gin.Context) {
	//	c.File("./frontend2/index.html") // Главная страница
	//})
	//router.GET("/auth/register", func(c *gin.Context) {
	//	c.File("./frontend2/register.html") // Страница регистрации
	//})
	//router.GET("/auth/login", func(c *gin.Context) {
	//	c.File("./frontend2/c.html") // Страница авторизации
	//})
	//router.GET("/chats", func(c *gin.Context) {
	//	c.File("./frontend2/chats.html") // Страница чатов
	//})
	publicPath := filepath.Join(".", "frontend2")

	// static
	router.StaticFile("/", filepath.Join(publicPath, "index.html"))
	router.StaticFile("/auth/register", filepath.Join(publicPath, "register.html"))
	router.StaticFile("/auth/login", filepath.Join(publicPath, "login.html"))
	router.StaticFile("/chats", filepath.Join(publicPath, "chats.html"))
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
