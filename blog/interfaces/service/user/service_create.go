package user

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
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

func (s *service) Create(ctx context.Context, req *CreateRequest) (id uint, err error) {
	user := mysql.User{}
	user.UserName = req.UserName
	user.UserNo = req.UserNo
	user.Addr = req.Addr
	user.Email = req.Email
	user.Password = req.Password
	t := uint(1)
	user.Level = &t

	id, err = user.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
