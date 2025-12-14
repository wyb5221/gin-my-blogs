package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 自定义Claims结构，添加你的业务字段
type MyClaims struct {
	Id                   uint   `json:"user_id"`
	UserNo               string `json:"username"`
	jwt.RegisteredClaims        // 标准Claims
}

func Test1() {
	// 用于签名的密钥（生产环境应从安全配置读取）
	var jwtSecret = []byte("your-secret-key")

	// 生成Token（使用自定义Claims）
	claims := MyClaims{
		Id:     1,
		UserNo: "Wyb",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "myapp",
		},
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("--token:", token)
	// 签名并获取Token字符串
	tokenString, err1 := token.SignedString(jwtSecret)
	if err1 != nil {
		fmt.Println("签名Token失败: %w", err1)
	}
	fmt.Println("生成的Token:", tokenString)

	// 解析Token
	parsedToken, err2 := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err2 != nil {
		fmt.Println("解析Token失败: %w", err2)
	}
	cla, ok := parsedToken.Claims.(*MyClaims)
	if ok && parsedToken.Valid {
		fmt.Println("--cla:", cla)
	} else {
		fmt.Println("无效的Token")
	}

}
