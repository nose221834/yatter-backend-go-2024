package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

// これは、絶対に変えたくない関数の形
type Status interface {
	Create(ctx context.Context, status *object.Status) error
	FindById(ctx context.Context, id int) (*object.Status, error)
}
