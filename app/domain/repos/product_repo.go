package repos

import "context"

type ProductRepo interface {
	Exist(ctx context.Context, productId []string) (bool, []string, error)
}
