package comment

import (
	"context"
	"fmt"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Content string `json:"content"`
	PostId  uint   `json:"postId"`
	UserId  uint   `json:"userId"`
}

func (s *service) AutoTable(ctx context.Context) {
	c := mysql.Comment{}
	c.AutoTable(s.db)
}

func (s *service) Create(ctx gin.Context, req *CreateRequest) (id uint, err error) {
	c := mysql.Comment{}
	c.Content = req.Content
	c.PostId = &req.PostId
	c.UserId = &req.UserId
	status := uint(1)
	c.Status = &status

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		c.UserId = &jwtClaims.Id
		fmt.Printf("用户ID: %d, 用户名: %s\n", jwtClaims.Id, jwtClaims.UserName)
	}

	id, err = c.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
