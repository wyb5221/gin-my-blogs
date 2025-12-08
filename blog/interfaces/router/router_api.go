package router

import (
	"gin-my-blogs/blog/interfaces/api/user"
	"gin-my-blogs/blog/interfaces/base"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterRoutes(r *gin.Engine) {
	db := base.InitDb()
	var logger *zap.Logger
	userHandler := user.New(logger, db)
	//创建分组
	gp := r.Group("/blogs")

	gp.POST("/user/add", userHandler.Create())

}
