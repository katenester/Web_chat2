package message

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/katenester/Web_chat2/backend/internal/models"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (t *MessagePostgres) GetMessage(userId int, friendId int) ([]models.Message, error) {
	return nil, errors.New("Not implemented")
}
func (t *MessagePostgres) Send(userId, friendId int, message models.Message) error {
	return errors.New("Not implemented")
}
