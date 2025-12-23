package user

import (
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

func (s *service) Delete(ctx gin.Context, id uint) (err error) {
	u := &mysql.User{}
	err = u.Delete(s.db, id)
	return err
}
