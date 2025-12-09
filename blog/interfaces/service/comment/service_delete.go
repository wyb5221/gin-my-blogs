package comment

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"
)

func (s *service) Delete(ctx context.Context, id uint) (err error) {
	c := &mysql.Comment{}
	err = c.Delete(s.db, id)
	return err
}
