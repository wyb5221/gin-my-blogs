package user

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

func (s *service) Delete(ctx context.Context, id uint) (err error) {
	u := &mysql.User{}
	err = u.Delete(s.db, id)
	return err
}
