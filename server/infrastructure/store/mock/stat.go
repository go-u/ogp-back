package mock

import (
	"context"
	domain "server/domain/model"
)

type StatStore struct {
	OnGet func(context.Context) ([]*domain.Stat, error)
}

func (s *StatStore) Get(ctx context.Context) ([]*domain.Stat, error) {
	return s.OnGet(ctx)
}
