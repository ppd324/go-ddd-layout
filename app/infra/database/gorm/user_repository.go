package gorm

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/domain/entity"
	"github.com/ppd324/go-ddd-layout/app/domain/repos"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repos.UserRepo {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetById(ctx context.Context, userId int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Save(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}
