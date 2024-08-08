package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

// これは、絶対に変えたくない関数の形
type Timeline interface {
	Public(ctx context.Context, limit int, since int) (*object.Status, error)
}
