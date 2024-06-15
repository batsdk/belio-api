package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `gorm:"not null" json:"email" gorm:"unique"`
	SubID    string `json:"sub_id" gorm:"unique"`
	Links    []Link `json:"links"`
}

type Link struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Link   string `json:"link" gorm:"not null"`
	UserID uint   `json:"user_id"`
}
