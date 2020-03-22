package usecase

import (
	"context"
	"server/domain/repository"
	pb "server/etc/protocol"
)

type BookmarkUsecase interface {
	Get(context.Context, uint64) ([]*pb.Stat, error)
	Add(context.Context, uint64, string) error
	Delete(context.Context, uint64, string) error
}

type bookmarkUsecase struct {
	Repo repository.BookmarkRepository
}

func NewBookmarkUsecase(bookmarkRepo repository.BookmarkRepository) BookmarkUsecase {
	bookmarkUsecase := bookmarkUsecase{bookmarkRepo}
	return &bookmarkUsecase
}

//func (u *bookmarkUsecase) Get(ctx context.Context, user *domain.User) ([]*domain.Bookmark, error) {
func (u *bookmarkUsecase) Get(ctx context.Context, userID uint64) ([]*pb.Stat, error) {
	statEntities, err := u.Repo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	if statEntities == nil {
		return make([]*pb.Stat, 0), nil
	}
	pbStats := convToStatsProto(statEntities)
	return pbStats, err
}

func (u *bookmarkUsecase) Add(ctx context.Context, userID uint64, fqdn string) error {
	err := u.Repo.Create(ctx, userID, fqdn)
	return err
}

func (u *bookmarkUsecase) Delete(ctx context.Context, userID uint64, fqdn string) error {
	err := u.Repo.Delete(ctx, userID, fqdn)
	return err
}
