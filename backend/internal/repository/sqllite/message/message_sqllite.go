package message

import (
	"database/sql"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"time"
)

type MessageSQLLite struct {
	db *sql.DB
}

func NewMessageSQLLite(db *sql.DB) *MessageSQLLite {
	return &MessageSQLLite{db: db}
}

func (t *MessageSQLLite) GetMessage(userId int, friendId int) ([]models.Message, error) {
	// SQL-запрос для получения всех сообщений из чата между двумя пользователями
	query := `
		SELECT m.id, m.chat_id, m.sender_id, m.message 
		FROM Messages m
		INNER JOIN Chats c ON m.chat_id = c.id
		WHERE (c.user1_id = ? AND c.user2_id = ?) 
		   OR (c.user1_id = ? AND c.user2_id = ?)
		ORDER BY m.sent_at ASC
	`

	// Выполнение запроса с подстановкой значений пользователей
	rows, err := t.db.Query(query, userId, friendId, friendId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	// Чтение строк результата запроса
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.Id, &msg.ChatId, &msg.SenderId, &msg.Message); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	// Проверка на наличие ошибок при сканировании строк
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
func (t *MessageSQLLite) Send(userId, friendId int, message models.Message) error {
	// Сначала нужно найти существующий чат между двумя пользователями
	var chatId int
	queryChat := `
		SELECT id 
		FROM Chats
		WHERE (user1_id = ? AND user2_id = ?) 
		   OR (user1_id = ? AND user2_id = ?)
	`
	err := t.db.QueryRow(queryChat, userId, friendId, friendId, userId).Scan(&chatId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Чат не существует, создаем новый
			createChatQuery := `INSERT INTO Chats (user1_id, user2_id) VALUES (?, ?)`
			result, err := t.db.Exec(createChatQuery, userId, friendId)
			if err != nil {
				return err
			}
			// Получаем ID нового чата
			chatId64, err := result.LastInsertId()
			if err != nil {
				return err
			}
			chatId = int(chatId64)
		} else {
			return err
		}
	}

	// Теперь можно вставить сообщение в таблицу Messages
	queryMessage := `
		INSERT INTO Messages (chat_id, sender_id, message, sent_at) 
		VALUES (?, ?, ?, ?)
	`
	_, err = t.db.Exec(queryMessage, chatId, userId, message.Message, time.Now())
	if err != nil {
		return err
	}

	return nil
}
