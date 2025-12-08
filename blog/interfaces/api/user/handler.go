package user

import (
	"gin-my-blogs/blog/interfaces/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	Create() func(ctx *gin.Context)
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
