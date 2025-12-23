package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string
	Content   string
	Type      string
	WorkCount *uint64
	Status    *uint
	UserId    *uint
	Comments  []Comment
}

func (p *Post) AutoTable(db *gorm.DB) {
	res := db.AutoMigrate(&Post{})
	if res != nil {
		fmt.Println("", res.Error())
	}
}

func (p *Post) Create(db *gorm.DB) (id uint, err error) {
	if err = db.Create(p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (p *Post) DetailById(db *gorm.DB, id uint) (post *Post, err error) {
	var pt = &Post{}
	result := db.First(pt, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return pt, nil
}

func (p *Post) List(db *gorm.DB) (posts *[]Post, err error) {
	var ps []Post
	query := db.Debug().Model(&ps)
	if p.Title != "" {
		query.Where("title like ?", "%"+p.Title+"%")
	}
	if p.Content != "" {
		query.Where("content like ?", "%"+p.Content+"%")
	}
	if p.Type != "" {
		query.Where("type = ?", p.Type)
	}
	workCount := p.WorkCount
	if *workCount != 0 {
		query.Where("work_count >= ?", workCount)
	}
	userId := p.UserId
	if *userId != 0 {
		query.Where("user_id = ?", userId)
	}
	return &ps, query.Find(&ps).Error
}

func (p *Post) Delete(db *gorm.DB, id uint) error {
	result := db.Delete(&Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Post) Updates(db *gorm.DB) error {
	// 创建更新结构体
	updateData := Post{}
	selectFields := []string{} // 记录要更新的字段
	// 动态判断并设置字段
	if p.Title != "" {
		updateData.Title = p.Title
		selectFields = append(selectFields, "title")
	}

	if p.Type != "" {
		updateData.Type = p.Type
		selectFields = append(selectFields, "type")
	}

	if p.Content != "" {
		updateData.Content = p.Content
		selectFields = append(selectFields, "content")
		l := uint64(len(p.Content))
		updateData.WorkCount = &l
		selectFields = append(selectFields, "work_count")
	}

	result := db.Model(&Post{}).Where("id=?", p.ID).Select(selectFields).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Post) GetPostCommentsById(db *gorm.DB, id uint) (post *Post, err error) {
	var pt = &Post{}
	result := db.Preload("Comments").First(pt, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return pt, nil
}
