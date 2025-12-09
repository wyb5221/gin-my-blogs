package post

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
	"unicode/utf8"
)

type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
	UserId  uint   `json:"userId"`
}

func (s *service) AutoTable(ctx context.Context) {
	p := mysql.Post{}
	p.AutoTable(s.db)
}

func (s *service) Create(ctx context.Context, req *CreateRequest) (id uint, err error) {
	p := mysql.Post{}
	p.Title = req.Title
	p.Content = req.Content
	p.Type = req.Type
	p.UserId = &req.UserId
	lenght := uint64(utf8.RuneCountInString(req.Content))
	p.WorkCount = &lenght
	status := uint(1)
	p.Status = &status

	id, err = p.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
