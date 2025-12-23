package comment

import (
	"gin-my-blogs/blog/interfaces/service/comment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	// AutoTable 建表
	AutoTable() func(ctx *gin.Context)
	// Create 新增
	Create() func(ctx *gin.Context)
	// Detail 评论信息
	Detail() func(ctx *gin.Context)
	// List 评论列表
	List() func(ctx *gin.Context)
	// Delete 删除评论
	Delete() func(ctx *gin.Context)
	// Updates 修改评论
	Updates() func(ctx *gin.Context)
	// UpdateStatus 修改评论状态
	UpdateStatus() func(ctx *gin.Context)
}

type handler struct {
	logger         *zap.Logger
	commentService comment.Service
}

func New(logger *zap.Logger, db *gorm.DB) Handler {
	return &handler{
		logger:         logger,
		commentService: comment.New(db),
	}
}
