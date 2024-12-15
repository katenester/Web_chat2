package auth

import (
	"database/sql"
	"errors"
	"github.com/katenester/Web_chat2/backend/internal/models"
)

type AuthSQLLite struct {
	db *sql.DB
}

func NewAuthSQLLite(db *sql.DB) *AuthSQLLite {
	return &AuthSQLLite{db}
}

func (a *AuthSQLLite) CreateUser(user models.User) error {
	return errors.New("Not implemented")
}
func (a *AuthSQLLite) GetUser(username, password string) (models.User, error) {
	return models.User{}, errors.New("Not implemented")
}
