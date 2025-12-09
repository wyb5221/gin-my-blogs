package user

import (
	"gin-my-blogs/blog/interfaces/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	// AutoTable 建表
	AutoTable() func(ctx *gin.Context)
	// Create 新增
	Create() func(ctx *gin.Context)
	// Detail 个人信息
	Detail() func(ctx *gin.Context)
	// List 用户列表
	List() func(ctx *gin.Context)
	// Delete 删除用户
	Delete() func(ctx *gin.Context)
	// ResetPassword 重置密码
	ResetPassword() func(ctx *gin.Context)
}

type handler struct {
	logger      *zap.Logger
	userService user.Service
}

func New(logger *zap.Logger, db *gorm.DB) Handler {
	return &handler{
		logger:      logger,
		userService: user.New(db),
	}
}
