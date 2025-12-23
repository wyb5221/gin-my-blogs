package comment

import (
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Content string `json:"content"`
	PostId  uint   `json:"postId"`
	UserId  uint   `json:"userId"`
}

func (s *service) DetailById(ctx gin.Context, id uint) (comment *mysql.Comment, err error) {
	c := &mysql.Comment{}
	i, err := c.DetailById(s.db, id)
	if err != nil {
		return nil, err
	}
	return i, err
}

func (s *service) List(ctx gin.Context, req *ListRequest) (comments *[]mysql.Comment, err error) {
	c := mysql.Comment{}
	c.Content = req.Content
	c.PostId = &req.PostId
	c.UserId = &req.UserId

	ps, err := c.List(s.db)
	if err != nil {
		return nil, err
	}
	return ps, err
}
