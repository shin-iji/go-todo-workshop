package database

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
