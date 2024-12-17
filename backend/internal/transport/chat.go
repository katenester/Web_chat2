package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"net/http"
)

func (h *Handler) getAllChats(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	// call service method
	chats, err := h.service.Chat.GetAll(UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Подготовим слайс для ответа
	var response []map[string]string

	// Проходим по каждому чату
	for _, chat := range chats {
		var otherUserId int

		// Определяем, какой пользователь является собеседником
		if chat.UserId == UserId {
			otherUserId = chat.User2Id
		} else {
			otherUserId = chat.UserId
		}

		// Получаем имя пользователя по id
		userName, err := h.service.Authorization.GetUserId(otherUserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve user name"})
			return
		}

		// Добавляем данные в слайс
		response = append(response, map[string]string{"user": userName})
	}

	// Возвращаем ответ в формате JSON
	c.JSON(http.StatusOK, response)
}

func (h *Handler) createChat(c *gin.Context) {
	//  Получаем userId из авторизации
	UserId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var userFriendId int
	// Извлекаем параметр user_name из URL
	userFriendId, err = h.service.Authorization.GetUserLogin(c.Param("user_name"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chats := models.Chat{UserId: UserId, User2Id: userFriendId}
	err = h.service.Chat.Create(chats)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Create chats successfully",
	})
}
