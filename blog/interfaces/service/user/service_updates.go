package user

import (
	"context"
	"fmt"
	"gin-my-blogs/blog/interfaces/mysql"
)

type UpdatesRequest struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	UserNo   string `json:"userNo"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *service) UpdatePassword(ctx context.Context, req *UpdatesRequest) (err error) {
	fmt.Println("---2---")
	user := mysql.User{}
	user.ID = req.Id
	user.Password = req.Password
	err = user.UpdatePassword(s.db)
	return err
}
