package post

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

func (s *service) Delete(ctx context.Context, id uint) (err error) {
	p := &mysql.Post{}
	err = p.Delete(s.db, id)
	return err
}
