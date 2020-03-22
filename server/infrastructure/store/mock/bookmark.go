package mock

import (
	"context"
	domain "server/domain/model"
)

type BookmarkStore struct {
	OnGet    func(context.Context, uint64) ([]*domain.Stat, error)
	OnCreate func(context.Context, uint64, string) error
	OnDelete func(context.Context, uint64, string) error
}

func (s *BookmarkStore) Get(ctx context.Context, userID uint64) ([]*domain.Stat, error) {
	return s.OnGet(ctx, userID)
}

func (s *BookmarkStore) Create(ctx context.Context, userID uint64, fqdn string) error {
	return s.OnCreate(ctx, userID, fqdn)
}

func (s *BookmarkStore) Delete(ctx context.Context, userID uint64, fqdn string) error {
	return s.OnDelete(ctx, userID, fqdn)
}
