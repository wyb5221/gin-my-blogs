package comment

import (
	"errors"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type UpdatesRequest struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
	PostId  uint   `json:"postId"`
}

func (s *service) Updates(ctx gin.Context, req *UpdatesRequest) (err error) {
	c := mysql.Comment{}
	c.ID = req.Id
	c.Content = req.Content

	comment, err := c.DetailById(s.db, req.Id)
	if err != nil {
		return err
	}

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != *comment.UserId {
			return errors.New("只能修改自己的评论")
		}
	}

	err = c.Updates(s.db)
	return err
}

func (s *service) UpdateStatus(ctx gin.Context, req *UpdatesRequest) (err error) {
	c := mysql.Comment{}
	c.ID = req.Id
	c.PostId = &req.PostId
	p := mysql.Post{}
	post, err := p.DetailById(s.db, req.PostId)
	if err != nil {
		return err
	}

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != *post.UserId {
			return errors.New("只能隐藏自己文章的评论")
		}
	}
	_, err = c.GetCommentByPostId(s.db)
	if err != nil {
		return errors.New("该条评论不属于这篇文章")
	}

	status := uint(9)
	c.Status = &status

	err = c.UpdateStatus(s.db)
	return err
}
