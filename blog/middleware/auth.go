package middleware

import (
	"gin-my-blogs/blog/common/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 从Header获取Token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求未携带Token，无权限访问",
			})
			ctx.Abort()
			return
		}
		// 2. 检查Token格式（Bearer token）
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token格式错误，格式应为: Bearer token",
			})
			ctx.Abort()
			return
		}

		// 3. 解析并验证Token
		tokenString := parts[1]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的Token: " + err.Error(),
			})
			ctx.Abort()
			return
		}
		// 4. Token验证成功，将用户信息存入上下文
		ctx.Set("userId", claims.Id)
		ctx.Set("userNo", claims.UserNo)
		ctx.Set("userName", claims.UserName)
		ctx.Set("claims", claims)

		// 5. 继续处理请求
		ctx.Next()
	}

}

// OptionalAuthMiddleware 可选认证中间件（登录和非登录用户都能访问）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenString := parts[1]
			if claims, err := jwt.ParseToken(tokenString); err == nil {
				// Token有效，设置用户信息
				c.Set("userId", claims.Id)
				c.Set("userNo", claims.UserNo)
				c.Set("userName", claims.UserName)
				c.Set("claims", claims)
			}
			// Token无效也不阻止请求，当作未登录用户处理
		}

		c.Next()
	}
}
