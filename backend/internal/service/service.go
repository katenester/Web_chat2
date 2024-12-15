package service

import (
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Chat interface {
	Create(userId int, chat models.Chat) error
	GetAll(userId int) ([]models.Chat, error)
}

type Message interface {
	GetMessage(userId int, friendId int) ([]models.Message, error)
	Send(userId, friendId int, message models.Message) error
}

type Service struct {
	Authorization
	Chat
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Chat:          NewTodoListService(repos.TodoList),
		Message:       NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
