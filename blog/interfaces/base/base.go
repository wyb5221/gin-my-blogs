package base

import (
	mysqls "gin-my-blogs/blog/interfaces/mysql"

	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * 创建数据库连接
 */
func InitDb() *gorm.DB {
	db, err :=
		gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&User{})
	return db
}

/**
 * 创建数据库表
 */
func CreateTable() {
	db := InitDb()
	db.AutoMigrate(&mysqls.User{})
	db.AutoMigrate(&mysqls.Post{})
	db.AutoMigrate(&mysqls.Comment{})
}
