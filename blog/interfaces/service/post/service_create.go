package post

import (
	"context"
	"fmt"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
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

func (s *service) Create(ctx gin.Context, req *CreateRequest) (id uint, err error) {
	p := mysql.Post{}
	p.Title = req.Title
	p.Content = req.Content
	p.Type = req.Type
	lenght := uint64(utf8.RuneCountInString(req.Content))
	p.WorkCount = &lenght
	status := uint(1)
	p.Status = &status

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		p.UserId = &jwtClaims.Id
		fmt.Printf("用户ID: %d, 用户名: %s\n", jwtClaims.Id, jwtClaims.UserName)
	}

	id, err = p.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
