package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	UsersTable    = "Users"
	ChatsTable    = "Chats"
	MessagesTable = "Messages"
)

//type Config struct {
//	Host     string
//	Port     string
//	Username string
//	Password string
//	DBName   string
//	SSLMode  string
//}

func NewSQLLite() (*sql.DB, error) {
	return sql.Open("sqlite3", "main.db")
}
