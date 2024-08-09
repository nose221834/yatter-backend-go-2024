package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

// *statusを返す→*statusに紐づいているCreateやFindも返している？
func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) Public(ctx context.Context, limit int) (*object.Timeline, error) {
	var timeline object.Timeline
	err := t.db.SelectContext(ctx, &timeline, "select * from status order by id desc limit ?", limit)
	if len(timeline) == 0 {
		return nil, fmt.Errorf("not found timeline from db")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find timeline from db %w", err)
	}
	return &timeline, nil
}
