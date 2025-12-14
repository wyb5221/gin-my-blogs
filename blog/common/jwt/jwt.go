package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 自定义Claims结构，添加你的业务字段
type MyClaims struct {
	Id                   uint   `json:"id"`
	UserNo               string `json:"userno"`
	UserName             string `json:"username"`
	jwt.RegisteredClaims        // 标准Claims
}

// 用于签名的密钥（生产环境应从安全配置读取）
var jwtSecret = []byte("your-secret-blogs")

// GenerateToken 生成JWT Token
func GenerateToken(id uint, userNo, userName string) (string, error) {
	// 生成Token（使用自定义Claims）
	claims := MyClaims{
		Id:       id,
		UserNo:   userNo,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "myapp",
		},
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名并获取Token字符串
	tokenString, err1 := token.SignedString(jwtSecret)
	if err1 != nil {
		return "获取Token失败！", errors.New("签名Token失败: %w")
	}
	return tokenString, nil
}

// ParseToken 解析并验证Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析Token
	parsedToken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	cla, ok := parsedToken.Claims.(*MyClaims)
	if ok && parsedToken.Valid {
		return cla, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
