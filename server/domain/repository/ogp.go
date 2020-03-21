package repository

import (
	"context"
	"server/domain/model"
)

type OgpRepository interface {
	Get(ctx context.Context, fqdn string) ([]*model.Ogp, error)
}
