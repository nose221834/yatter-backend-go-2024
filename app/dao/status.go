package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

// (a* status) →　登録する相手　→statusの一部として振る舞う
func (s *status) Create(ctx context.Context, status *object.Status) error {
	_, err := s.db.Exec("insert into status (AccountID,Content,CreatedAt) values (?,?,?)",
		status.AccountID, status.Content, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}
	return nil
}
