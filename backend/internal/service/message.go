package service

import (
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewTodoItemService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (m *MessageService) GetMessage(userId int, friendId int) ([]models.Message, error) {
	return m.repo.GetMessage(userId, friendId)
}
func (m *MessageService) Send(userId, friendId int, message models.Message) error {
	return m.repo.Send(userId, friendId, message)
}
