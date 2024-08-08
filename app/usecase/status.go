package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, content string, account_id int) (*CreateStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

type FindStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context, content string, account_id int) (*CreateStatusDTO, error) {
	status := object.NewStatus(content, account_id)

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
	// domainに依存させることで、保守性を上げている
	if err := s.statusRepo.Create(ctx, status); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: status,
	}, nil
}

func (s *status) FindById(ctx context.Context, id int) (*FindStatusDTO, error) {
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

	status, err := s.statusRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &FindStatusDTO{
		Status: status,
	}, nil
}
