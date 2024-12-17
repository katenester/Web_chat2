package models

type Message struct {
	Id       int    `json:"-" db:"id"`
	ChatId   int    `json:"-" db:"chat_id"`
	SenderId int    `json:"sender"  db:"sender_id"`
	Message  string `json:"message" binding:"required" db:"message"`
}
