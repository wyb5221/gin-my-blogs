package user

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx context.Context, adminData *CreateRequest) (id uint, err error)
}
type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}
