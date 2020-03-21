package repository

import (
	"context"
	"server/domain/model"
)

type UserRepository interface {
	GetByUID(ctx context.Context, uid string) (*model.User, error)
	Create(ctx context.Context, uid string, name string) error
	Delete(ctx context.Context, uid string) error
}
