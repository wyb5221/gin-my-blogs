package post

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type ListRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	UserId    uint   `json:"userId"`
	WorkCount uint64 `json:"workCount"`
}

func (s *service) DetailById(ctx context.Context, id uint) (user *mysql.Post, err error) {
	p := &mysql.Post{}
	i, err := p.DetailById(s.db, id)
	if err != nil {
		return nil, err
	}
	return i, err
}

func (s *service) List(ctx context.Context, req *ListRequest) (posts *[]mysql.Post, err error) {
	p := mysql.Post{}
	p.Title = req.Title
	p.Content = req.Content
	p.Type = req.Type
	p.UserId = &req.UserId
	p.WorkCount = &req.WorkCount

	ps, err := p.List(s.db)
	if err != nil {
		return nil, err
	}
	return ps, err
}
