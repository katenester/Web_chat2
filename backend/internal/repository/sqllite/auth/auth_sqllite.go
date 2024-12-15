package auth

import (
	"database/sql"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository/sqllite/config"
	"time"
)

type AuthSQLLite struct {
	db *sql.DB
}

func NewAuthSQLLite(db *sql.DB) *AuthSQLLite {
	return &AuthSQLLite{db}
}

func (a *AuthSQLLite) CreateUser(user models.User) error {
	// Вставляем нового пользователя в базу данных
	query := `INSERT INTO ` + config.UsersTable + ` (username, password, created_at) VALUES (?, ?, ?)`
	_, err := a.db.Exec(query, user.Username, user.Password, time.Now())
	// Если возникает ошибка (например уникальности) (ошибка вставки), возвращаем кастомную ошибку
	if err != nil {
		return err
	}
	return nil
}
func (a *AuthSQLLite) GetUser(username, password string) (models.User, error) {
	var user models.User
	// Запрос для поиска пользователя по имени
	query := `SELECT id, username, password FROM ` + config.UsersTable + ` WHERE username = ?`
	row := a.db.QueryRow(query, username)
	// Извлечение данных пользователя
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
