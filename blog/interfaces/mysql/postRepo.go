package mysql

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string
	Content   string
	Type      string
	WorkCount uint64
	UserId    uint
	Comments  []Comment
}
