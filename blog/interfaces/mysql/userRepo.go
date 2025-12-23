package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	UserNo   string
	Addr     string
	Password string
	Email    string
	Level    *uint
	CreateBy *uint
	BlogNum  *uint
	Posts    []Post
}

func (u *User) AutoTable(db *gorm.DB) {
	res := db.AutoMigrate(&User{})
	if res != nil {
		fmt.Println("", res.Error())
	}
}

func (u *User) Create(db *gorm.DB) (id uint, err error) {
	var user User
	result := db.Where("user_no = ?", u.UserNo).First(&user)

	// 通过 RowsAffected 判断是否找到记录
	if result.RowsAffected > 0 {
		// 用户已存在，返回错误
		return 0, fmt.Errorf("用户已存在，用户编号: %s", u.UserNo)
	}
	if err = db.Create(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (u *User) DetailById(db *gorm.DB, id uint) (user *User, err error) {
	var ur = &User{}
	result := db.First(ur, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ur, nil
}

func (u *User) QueryByUserNoAndPwd(db *gorm.DB) (user *User, err error) {
	var ur = &User{}
	result := db.Where("user_no = ?  ", u.UserNo).First(ur)
	if result.Error != nil {
		return nil, result.Error
	}
	return ur, nil
}

func (u *User) List(db *gorm.DB) (user *[]User, err error) {
	var users []User
	query := db.Model(&User{})
	if u.UserName != "" {
		query.Where("user_name like ?", "%"+u.UserName+"%")
	}
	if u.UserNo != "" {
		query.Where("user_no = ?", u.UserNo)
	}
	if u.Email != "" {
		query.Where("email = ?", u.Email)
	}
	if u.Addr != "" {
		query.Where("addr like ?", "%"+u.Addr+"%")
	}
	level := u.Level
	if *level != 0 {
		query.Where("level = ?", level)
	}
	return &users, query.Find(&users).Error
}

func (u *User) Delete(db *gorm.DB, id uint) error {
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) UpdatePassword(db *gorm.DB) error {
	fmt.Println("---3---")
	result := db.Model(&User{}).Where("id=?", u.ID).Update("password", u.Password)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
