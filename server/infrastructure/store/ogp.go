package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/queries/qm"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"time"
)

func NewOgpRepository(sqlHandler SqlHandler) repository.OgpRepository {
	ogpRepository := OgpStore{sqlHandler}
	return &ogpRepository
}

type OgpStore struct {
	SqlHandler
}

func (s *OgpStore) Get(ctx context.Context, fqdn string) ([]*domain.Ogp, error) {
	recent := time.Now().AddDate(0, 0, -7)
	limit := 10
	ogps, err := models.Ogps(models.OgpWhere.FQDN.EQ(fqdn), models.OgpWhere.Date.GTE(recent), qm.OrderBy("-"+models.OgpColumns.Date), qm.Limit(limit)).All(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}

	var ogpEntities []*domain.Ogp
	for _, ogp := range ogps {
		ogpEntity := convToOgpEntity(ogp)
		ogpEntities = append(ogpEntities, ogpEntity)
	}
	return ogpEntities, nil
}

func convToOgpEntity(ogp *models.Ogp) *domain.Ogp {
	ogpEntity := &domain.Ogp{
		Date:    ogp.Date,
		FQDN:    ogp.FQDN,
		Host:    ogp.Host,
		TweetID: ogp.TweetID,
		Type:    ogp.Type,
		Lang:    ogp.Lang,
	}
	return ogpEntity
}
