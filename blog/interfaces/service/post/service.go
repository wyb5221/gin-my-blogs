package post

import (
	"context"

	"gin-my-blogs/blog/interfaces/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	AutoTable(ctx context.Context)

	Create(ctx gin.Context, req *CreateRequest) (id uint, err error)

	DetailById(ctx gin.Context, id uint) (user *mysql.Post, err error)

	List(ctx gin.Context, req *ListRequest) (user *[]mysql.Post, err error)

	Updates(ctx gin.Context, req *UpdatesRequest) (err error)

	Delete(ctx gin.Context, id uint) (err error)

	GetPostCommentsById(ctx gin.Context, id uint) (post *mysql.Post, err error)
}
type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
