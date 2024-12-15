package service

import (
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewTodoListService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Create(userId int, chat models.Chat) error {
	return s.repo.Create(userId, chat)
}
func (s *ChatService) GetAll(userId int) ([]models.Chat, error) {
	return s.repo.GetAll(userId)
}
