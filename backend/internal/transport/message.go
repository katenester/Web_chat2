package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"net/http"
)

func (h *Handler) getAllMessage(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	// Извлекаем параметр user_name из URL
	userFriendId, err := h.service.Authorization.GetUserLogin(c.Param("user_name"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var messages []models.Message
	messages, err = h.service.Message.GetMessage(UserId, userFriendId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Подготовим слайс для ответа
	var response []map[string]string
	for _, message := range messages {
		var sender string
		sender, err := h.service.Authorization.GetUserId(message.SenderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response = append(response, map[string]string{"sender": sender, "message": message.Message})
	}
	// Возвращаем ответ в формате JSON
	c.JSON(http.StatusOK, response)
}
func (h *Handler) sendMessage(c *gin.Context) {
	//  Take value UserId from context
	UserId, ok := getUserId(c)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	// Извлекаем параметр user_name из URL
	userFriendId, err := h.service.Authorization.GetUserLogin(c.Param("user_name"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Извлекаем сообщение из тела
	var input models.Message
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.Message.Send(UserId, userFriendId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}
