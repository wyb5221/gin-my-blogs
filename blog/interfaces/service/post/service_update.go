package post

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type UpdatesRequest struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (s *service) Updates(ctx context.Context, req *UpdatesRequest) (err error) {
	p := mysql.Post{}
	p.ID = req.Id
	p.Title = req.Title
	p.Content = req.Content
	p.Type = req.Type
	err = p.Updates(s.db)
	return err
}
