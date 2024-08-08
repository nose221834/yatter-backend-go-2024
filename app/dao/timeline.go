package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	rows, err := t.db.QueryxContext(ctx, "select * from status order by id desc limit ?", limit)
	var timeline object.Timeline

	defer rows.Close()

	for rows.Next() {
		var status object.Status
		err := rows.StructScan(&status)
		if err != nil {
			log.Fatalln(err)
		}
		timeline = append(timeline, status)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find timeline from db %w", err)
	}
	return &timeline, nil
}
