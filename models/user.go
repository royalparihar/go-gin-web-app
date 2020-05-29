package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	IsAdmin   string    `json:"role"`
	UserName  string    `json: "user_name"`
	Password  string    `json: "password"`
	IsDelete  bool      `json: "is_delete"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
