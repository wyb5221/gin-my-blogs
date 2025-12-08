package mysql

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserId  uint
	PostId  uint
}
