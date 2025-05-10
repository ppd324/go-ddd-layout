package repos

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/domain/entity"
)

type UserRepo interface {
	GetById(ctx context.Context, userId int) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
