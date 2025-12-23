package user

import (
	"errors"
	"fmt"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdatesRequest struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	UserNo   string `json:"userNo"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *service) UpdatePassword(ctx gin.Context, req *UpdatesRequest) (err error) {
	fmt.Println("---2---")
	user := mysql.User{}
	user.ID = req.Id

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// 或者获取完整的claims
	if claims, exists := ctx.Get("claims"); exists {
		jwtClaims := claims.(*jwt.MyClaims)
		tId := jwtClaims.Id
		if tId != req.Id {
			return errors.New("只能修改自己的密码")
		}
		// user.CreateBy = &jwtClaims.Id
		// fmt.Printf("用户ID: %d, 用户名: %s\n", jwtClaims.Id, jwtClaims.UserName)
	}

	err = user.UpdatePassword(s.db)
	return err
}
