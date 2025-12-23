package post

import (
	"errors"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
)

func (s *service) Delete(ctx gin.Context, id uint) (err error) {
	p := &mysql.Post{}
	post, err := p.DetailById(s.db, id)
	if err != nil {
		return err
	}

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != *post.UserId {
			return errors.New("只能删除自己的文章")
		}
	}

	err = p.Delete(s.db, id)
	return err
}
