package models

type Chat struct {
	Id      int `json:"-" db:"id"`
	UserId  int `json:"user1_id" binding:"required" db:"user1_id"`
	User2Id int `json:"user2_id" binding:"required" db:"user2_id"`
}
