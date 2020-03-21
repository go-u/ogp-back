package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/queries/qm"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"time"
)

func NewStatRepository(sqlHandler SqlHandler) repository.StatRepository {
	statRepository := StatStore{sqlHandler}
	return &statRepository
}

type StatStore struct {
	SqlHandler
}

func (s *StatStore) Get(ctx context.Context) ([]*domain.Stat, error) {
	recent := time.Now().AddDate(0, 0, -7)

	// todo: configurable value
	minimum := 20
	stats, err := models.Stats(models.StatWhere.Date.GTE(recent), models.StatWhere.Count.GTE(uint(minimum)), qm.Limit(10000)).All(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	sumStats := sumStats(stats)

	var statEntities []*domain.Stat
	for _, stat := range sumStats {
		statEntity := convToStatEntity(stat)
		statEntities = append(statEntities, statEntity)
	}
	return statEntities, err
}

func convToStatEntity(stat *models.Stat) *domain.Stat {
	statEntity := &domain.Stat{
		Date:        stat.Date,
		FQDN:        stat.FQDN,
		Host:        stat.Host,
		Count:       stat.Count,
		Title:       stat.Title,
		Description: stat.Description,
		Image:       stat.Image,
		Type:        stat.Type,
		Lang:        stat.Lang,
	}
	return statEntity
}

func sumStats(stats []*models.Stat) []*models.Stat {
	// 集計用のFQDNをキーとしたMap
	statMap := map[string]*models.Stat{}
	for _, stat := range stats {
		v, exist := statMap[stat.FQDN]
		if exist {
			stat.Count += v.Count
		}
		statMap[stat.FQDN] = stat
	}
	var sumStats = []*models.Stat{}
	for _, v := range statMap {
		sumStats = append(sumStats, v)
	}
	return sumStats
}
