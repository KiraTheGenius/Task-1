package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `db:"user_name" json:"user_name"`
	Password  string `db:"password" json:"password"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Phone     string `gorm:"size:10"`
	Email     string `gorm:"size:255"`
}
