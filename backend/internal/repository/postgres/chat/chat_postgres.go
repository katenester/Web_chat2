package chat

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/katenester/Web_chat2/backend/internal/models"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (t *ChatPostgres) Create(userId int, chat models.Chat) error {
	return errors.New("Not implemented")
}
func (t *ChatPostgres) GetAll(userId int) ([]models.Chat, error) {
	return nil, errors.New("Not implemented")
}
