package user

import (
	"context"
	"fmt"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	UserName string `json:"userName"`
	UserNo   string `json:"userNo"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *service) AutoTable(ctx context.Context) {
	user := mysql.User{}
	user.AutoTable(s.db)
}

func (s *service) Create(ctx gin.Context, req *CreateRequest) (id uint, err error) {
	user := mysql.User{}
	user.UserName = req.UserName
	user.UserNo = req.UserNo
	user.Addr = req.Addr
	user.Email = req.Email
	user.Password = req.Password
	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		user.CreateBy = &jwtClaims.Id
		fmt.Printf("用户ID: %d, 用户名: %s\n", jwtClaims.Id, jwtClaims.UserName)
	}

	t := uint(1)
	user.Level = &t

	id, err = user.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
