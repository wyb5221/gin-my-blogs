package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	UserNO   string
	Addr     string
	Password string
	Email    string
	Level    uint
	Posts    []Post
}
