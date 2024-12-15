package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/katenester/Web_chat2/backend/internal/models"
	"github.com/katenester/Web_chat2/backend/internal/repository/postgres/config"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a *AuthPostgres) CreateUser(user models.User) error {
	// Вставляем нового пользователя в базу данных
	query := `INSERT INTO ` + config.UsersTable + ` (username, password) VALUES (?, ?)`
	_, err := a.db.Exec(query, user.Username, user.Password)
	// Если возникает ошибка (например уникальности) (ошибка вставки), возвращаем кастомную ошибку
	if err != nil {
		return err
	}
	return nil
}
func (a *AuthPostgres) GetUser(username, password string) (models.User, error) {
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
