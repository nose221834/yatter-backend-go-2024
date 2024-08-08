package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, content string, id int) (*AddStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type AddStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) AddStatus(ctx context.Context, content string, id int) (*AddStatusDTO, error) {
	status := object.NewStatus(content, id)

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	// domainを叩いている？daoには直接アクセスしていないらしい
	if err := s.statusRepo.AddStatus(ctx, status); err != nil {
		return nil, err
	}

	return &AddStatusDTO{
		Status: status,
	}, nil
}
