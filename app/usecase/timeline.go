package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	Public(ctx context.Context, limit int) (*PublicTimelineDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	timelineRepo repository.Timeline
}

type PublicTimelineDTO struct {
	Timeline *object.Timeline
}

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) Public(ctx context.Context, limit int) (*PublicTimelineDTO, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	timeline, err := t.timelineRepo.Public(ctx, limit)

	if err != nil {
		return nil, err
	}

	return &PublicTimelineDTO{
		Timeline: timeline,
	}, nil

}
