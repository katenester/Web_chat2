package message

import (
	"database/sql"
	"errors"
	"github.com/katenester/Web_chat2/backend/internal/models"
)

type MessageSQLLite struct {
	db *sql.DB
}

func NewMessageSQLLite(db *sql.DB) *MessageSQLLite {
	return &MessageSQLLite{db: db}
}

func (t *MessageSQLLite) GetMessage(userId int, friendId int) ([]models.Message, error) {
	return nil, errors.New("Not implemented")
}
func (t *MessageSQLLite) Send(userId, friendId int, message models.Message) error {
	return errors.New("Not implemented")
}
