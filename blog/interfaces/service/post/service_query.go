package post

import (
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	UserId    uint   `json:"userId"`
	WorkCount uint64 `json:"workCount"`
}

func (s *service) DetailById(ctx gin.Context, id uint) (post *mysql.Post, err error) {
	p := &mysql.Post{}
	i, err := p.DetailById(s.db, id)
	if err != nil {
		return nil, err
	}
	return i, err
}

func (s *service) List(ctx gin.Context, req *ListRequest) (posts *[]mysql.Post, err error) {
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

func (s *service) GetPostCommentsById(ctx gin.Context, id uint) (post *mysql.Post, err error) {
	p := &mysql.Post{}
	i, err := p.GetPostCommentsById(s.db, id)
	if err != nil {
		return nil, err
	}
	return i, err
}
