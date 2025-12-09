package comment

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type CreateRequest struct {
	Content string `json:"content"`
	PostId  uint   `json:"postId"`
	UserId  uint   `json:"userId"`
}

func (s *service) Create(ctx context.Context, req *CreateRequest) (id uint, err error) {
	c := mysql.Comment{}
	c.Content = req.Content
	c.PostId = &req.PostId
	c.UserId = &req.UserId

	id, err = c.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
