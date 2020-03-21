package mock

import (
	"context"
	domain "server/domain/model"
)

type OgpStore struct {
	OnGet func(context.Context, string) ([]*domain.Ogp, error)
}

func (s *OgpStore) Get(ctx context.Context, fqdn string) ([]*domain.Ogp, error) {
	return s.OnGet(ctx, fqdn)
}
