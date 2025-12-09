package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	Status  *uint
	UserId  *uint
	PostId  *uint
}

func (c *Comment) AutoTable(db *gorm.DB) {
	res := db.AutoMigrate(&Comment{})
	if res != nil {
		fmt.Println("", res.Error())
	}
}

func (c *Comment) Create(db *gorm.DB) (id uint, err error) {
	if err = db.Create(c).Error; err != nil {
		return 0, err
	}
	return c.ID, nil
}

func (c *Comment) DetailById(db *gorm.DB, id uint) (comments *Comment, err error) {
	var cs = &Comment{}
	result := db.First(cs, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return cs, nil
}

func (c *Comment) List(db *gorm.DB) (comments *[]Comment, err error) {
	var cs []Comment
	query := db.Model(&Comment{})
	if c.Content != "" {
		query.Where("content like ?", "%"+c.Content+"%")
	}
	userId := c.UserId
	if *userId != 0 {
		query.Where("user_id = ?", userId)
	}
	postId := c.PostId
	if *postId != 0 {
		query.Where("post_id = ?", postId)
	}
	return &cs, query.Find(&cs).Error
}

func (c *Comment) Delete(db *gorm.DB, id uint) error {
	result := db.Delete(&Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Comment) Updates(db *gorm.DB) error {
	result := db.Model(&Comment{}).Where("id=?", c.ID).Update("content", c.Content)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
