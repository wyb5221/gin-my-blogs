package post

import (
	"gin-my-blogs/blog/interfaces/service/post"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	// Create 新增
	Create() func(ctx *gin.Context)
	// Detail 文章信息
	Detail() func(ctx *gin.Context)
	// List 文章列表
	List() func(ctx *gin.Context)
	// Delete 删除文章
	Delete() func(ctx *gin.Context)
	// Updates 修改文章
	Updates() func(ctx *gin.Context)
}

type handler struct {
	logger      *zap.Logger
	postService post.Service
}

func New(logger *zap.Logger, db *gorm.DB) Handler {
	return &handler{
		logger:      logger,
		postService: post.New(db),
	}
}
