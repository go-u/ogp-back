package repository

import (
	"context"
	"server/domain/model"
)

type StatRepository interface {
	Get(ctx context.Context) ([]*model.Stat, error)
}
