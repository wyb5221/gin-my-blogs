package user

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type CreateRequest struct {
	UserName string
	UserNo   string
	Addr     string
	Password string
	Email    string
}

func (s *service) Create(ctx context.Context, req *CreateRequest) (id uint, err error) {
	user := mysql.User{}
	user.UserName = req.UserName
	user.UserNo = req.UserNo
	user.Addr = req.Addr
	user.Password = req.Password
	user.Level = 1

	id, err = user.Create(s.db)
	if err != nil {
		return 0, err
	}
	return
}
