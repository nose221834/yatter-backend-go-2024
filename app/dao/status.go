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

func (s *status) Find(ctx context.Context, id int) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db %w", err)
	}
	return entity, nil
}
