package post

import (
	"errors"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type UpdatesRequest struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (s *service) Updates(ctx gin.Context, req *UpdatesRequest) (err error) {
	p := mysql.Post{}
	p.ID = req.Id
	p.Title = req.Title
	p.Content = req.Content
	p.Type = req.Type

	post, err := p.DetailById(s.db, req.Id)
	if err != nil {
		return err
	}

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != *post.UserId {
			return errors.New("只能修改自己的文章")
		}
	}

	err = p.Updates(s.db)
	return err
}
