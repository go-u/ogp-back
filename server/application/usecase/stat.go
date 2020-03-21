package usecase

import (
	"context"
	domain "server/domain/model"
	"server/domain/repository"
	pb "server/etc/protocol"
)

type StatUsecase interface {
	Get(context.Context) ([]*pb.Stat, error)
}

type statUsecase struct {
	Repo repository.StatRepository
}

func NewStatUsecase(statRepo repository.StatRepository) StatUsecase {
	statUsecase := statUsecase{statRepo}
	return &statUsecase
}

func (u *statUsecase) Get(ctx context.Context) ([]*pb.Stat, error) {
	statEntities, err := u.Repo.Get(ctx)
	if err != nil {
		return nil, err
	}
	stats := convToStatsProto(statEntities)
	return stats, err
}

func convToStatsProto(statEntities []*domain.Stat) []*pb.Stat {
	var statsProto []*pb.Stat
	for _, statEntity := range statEntities {
		statProto := &pb.Stat{
			Fqdn:        statEntity.FQDN,
			Count:       int32(statEntity.Count),
			Title:       statEntity.Title,
			Description: statEntity.Description,
			Image:       statEntity.Image,
			Lang:        statEntity.Lang,
		}
		statsProto = append(statsProto, statProto)
	}
	return statsProto
}
