package models

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	UserName  string    `json:"username" binding:"required"`
	PassWord  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
