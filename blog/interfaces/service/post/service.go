package post

import (
	"context"

	"gin-my-blogs/blog/interfaces/mysql"

	"gorm.io/gorm"
)

type Service interface {
	AutoTable(ctx context.Context)

	Create(ctx context.Context, req *CreateRequest) (id uint, err error)

	DetailById(ctx context.Context, id uint) (user *mysql.Post, err error)

	List(ctx context.Context, req *ListRequest) (user *[]mysql.Post, err error)

	Updates(ctx context.Context, req *UpdatesRequest) (err error)

	Delete(ctx context.Context, id uint) (err error)
}
type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
