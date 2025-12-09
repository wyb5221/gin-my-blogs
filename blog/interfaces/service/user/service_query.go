package user

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

type ListRequest struct {
	UserName string `json:"userName"`
	UserNo   string `json:"userNo"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *service) DetailById(ctx context.Context, id uint) (user *mysql.User, err error) {
	u := &mysql.User{}
	i, err := u.DetailById(s.db, id)
	if err != nil {
		return nil, err
	}
	return i, err
}

func (s *service) List(ctx context.Context, req *ListRequest) (users *[]mysql.User, err error) {
	user := mysql.User{}
	user.UserName = req.UserName
	user.UserNo = req.UserNo
	user.Addr = req.Addr
	user.Password = req.Password

	us, err := user.List(s.db)
	if err != nil {
		return nil, err
	}
	return us, err
}
