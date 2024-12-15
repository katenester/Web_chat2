package chat

import (
	"database/sql"
	"errors"
	"github.com/katenester/Web_chat2/backend/internal/models"
)

type ChatSQLLite struct {
	db *sql.DB
}

func NewChatSQLLite(db *sql.DB) *ChatSQLLite {
	return &ChatSQLLite{db: db}
}

func (t *ChatSQLLite) Create(userId int, chat models.Chat) error {
	return errors.New("Not implemented")
}
func (t *ChatSQLLite) GetAll(userId int) ([]models.Chat, error) {
	return nil, errors.New("Not implemented")
}
