package comment

import (
	"errors"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

func (s *service) Delete(ctx gin.Context, id uint) (err error) {
	c := &mysql.Comment{}

	comment, err := c.DetailById(s.db, id)
	if err != nil {
		return err
	}

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != *comment.UserId {
			return errors.New("只能删除自己发的评论")
		}
	}

	err = c.Delete(s.db, id)
	return err
}
