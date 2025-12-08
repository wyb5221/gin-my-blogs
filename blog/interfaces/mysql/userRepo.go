package mysql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	UserNo   string
	Addr     string
	Password string
	Email    string
	Level    uint
	Posts    []Post
}

func (u *User) Create(db *gorm.DB) (id uint, err error) {
	if err = db.Create(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}
