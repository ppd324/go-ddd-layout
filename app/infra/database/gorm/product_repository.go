package gorm

import (
	"context"
	"github.com/ppd324/go-ddd-layout/app/domain/repos"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repos.ProductRepo {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Exist(ctx context.Context, productId []string) (bool, []string, error) {
	//TODO implement me
	panic("implement me")
}
