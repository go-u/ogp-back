package mock

import (
	"context"
	domain "server/domain/model"
)

type UserStore struct {
	OnGetByUID func(context.Context, string) (*domain.User, error)
	OnCreate   func(context.Context, string, string) error
	OnDelete   func(context.Context, string) error
}

func (s *UserStore) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	return s.OnGetByUID(ctx, uid)
}

func (s *UserStore) Create(ctx context.Context, uid string, name string) error {
	return s.OnCreate(ctx, uid, name)
}

func (s *UserStore) Delete(ctx context.Context, uid string) error {
	return s.OnDelete(ctx, uid)
}
