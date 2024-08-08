package dao

import (
	"context"
	"database/sql"
	"errors"
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

var _ repository.Status = (*status)(nil)

// *statusを返す→*statusに紐づいているCreateやFindも返している？
func NewTimeline(db *sqlx.DB) *status {
	return &status{db: db}
}

func (t *timeline) Public(ctx context.Context, id int) (*object.Timeline, error) {
	entity := new(object.Timeline)
	err := t.db.QueryRowxContext(ctx, "select * from status order by id desc limit ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db %w", err)
	}
	return entity, nil
}
