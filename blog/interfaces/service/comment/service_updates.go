package comment

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type UpdatesRequest struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
}

func (s *service) Updates(ctx context.Context, req *UpdatesRequest) (err error) {
	c := mysql.Comment{}
	c.ID = req.Id
	c.Content = req.Content
	err = c.Updates(s.db)
	return err
}
