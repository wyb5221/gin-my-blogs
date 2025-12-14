package user

import (
	"context"
	"errors"
	"gin-my-blogs/blog/common/jwt"
	"gin-my-blogs/blog/interfaces/mysql"
	// "github.com/golang-jwt/jwt/v4"
)

type ListRequest struct {
	UserName string `json:"userName"`
	UserNo   string `json:"userNo"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Level    uint   `json:"level"`
}

func (s *service) Login(ctx context.Context, req *ListRequest) (str string, err error) {
	user := mysql.User{}
	user.UserNo = req.UserNo
	user.Password = req.Password

	ue, err := user.QueryByUserNoAndPwd(s.db)
	if err != nil {
		return "用户账户或密码错误！", err
	}
	// 检查返回的对象是否为 nil
	if ue == nil {
		return "用户账户或密码错误！", errors.New("用户不存在")
	}

	token, err := jwt.GenerateToken(ue.ID, ue.UserNo, ue.UserName)

	return token, err

	// // 生成Token（使用自定义Claims）
	// claims := MyClaims{
	// 	Id:       ue.ID,
	// 	UserNo:   ue.UserNo,
	// 	UserName: ue.UserName,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 		Issuer:    "myapp",
	// 	},
	// }
	// // 生成 JWT
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// // 签名并获取Token字符串
	// tokenString, err1 := token.SignedString(jwtSecret)
	// if err1 != nil {
	// 	return "用户账户或密码错误！获取Token失败！", errors.New("签名Token失败: %w")
	// }

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
	user.Level = &req.Level
	user.Password = req.Password

	us, err := user.List(s.db)
	if err != nil {
		return nil, err
	}
	return us, err
}
