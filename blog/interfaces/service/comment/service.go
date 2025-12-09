package comment

import (
	"context"
	"gin-my-blogs/blog/interfaces/mysql"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) (id uint, err error)

	DetailById(ctx context.Context, id uint) (user *mysql.Comment, err error)

	List(ctx context.Context, req *ListRequest) (user *[]mysql.Comment, err error)

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
