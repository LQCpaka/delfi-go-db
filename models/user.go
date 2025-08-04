package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
}
