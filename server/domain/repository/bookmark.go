package repository

import (
	"context"
	"server/domain/model"
)

type BookmarkRepository interface {
	Get(ctx context.Context, userID uint64) ([]*model.Bookmark, error)
	Create(ctx context.Context, userID uint64, fqdn string) error
	Delete(ctx context.Context, userID uint64, fqdn string) error
}
