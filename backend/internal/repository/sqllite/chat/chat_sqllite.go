package chat

import (
	"database/sql"
	"fmt"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"time"
)

type ChatSQLLite struct {
	db *sql.DB
}

func NewChatSQLLite(db *sql.DB) *ChatSQLLite {
	return &ChatSQLLite{db: db}
}

func (t *ChatSQLLite) Create(chat models.Chat) error {
	// Вставляем новый чат в базу данных
	query := `INSERT INTO Chats (user1_id, user2_id, created_at) VALUES (?, ?, ?)`
	_, err := t.db.Exec(query, chat.UserId, chat.User2Id, time.Now())
	if err != nil {
		// Если произошла ошибка, например нарушение уникального ограничения, возвращаем её
		return fmt.Errorf("error creating chat: %v", err)
	}
	return nil
}
func (t *ChatSQLLite) GetAll(userId int) ([]models.Chat, error) {
	// Запрос для получения всех чатов, где пользователь является либо user1_id, либо user2_id
	query := `SELECT id, user1_id, user2_id FROM Chats WHERE user1_id = ? OR user2_id = ?`

	rows, err := t.db.Query(query, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat

	// Проходим по всем полученным строкам и заполняем слайс чатов
	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.Id, &chat.UserId, &chat.User2Id); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	// Проверка на ошибки при сканировании строк
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}
