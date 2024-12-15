package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository/postgres/auth"
	"github.com/katenester/Web_chat2/backend/internal/repository/postgres/chat"
	"github.com/katenester/Web_chat2/backend/internal/repository/postgres/message"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(username, password string) (models.User, error)
}

type Chat interface {
	Create(userId int, chat models.Chat) error
	GetAll(userId int) ([]models.Chat, error)
}

type Message interface {
	GetMessage(userId int, friendId int) ([]models.Message, error)
	Send(userId, friendId int, message models.Message) error
}

type Repository struct {
	Authorization
	Chat
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(db),
		Chat:          chat.NewChatPostgres(db),
		Message:       message.NewMessagePostgres(db),
	}
}
