package usecase

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	domain "server/domain/model"
	"server/domain/repository"
	pb "server/etc/protocol"
)

type BookmarkUsecase interface {
	Get(context.Context, uint64) ([]*pb.Bookmark, error)
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
func (u *bookmarkUsecase) Get(ctx context.Context, userID uint64) ([]*pb.Bookmark, error) {
	bookmarkEntities, err := u.Repo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	if bookmarkEntities == nil {
		return make([]*pb.Bookmark, 0), nil
	}
	bookmarks, err := convToBookmarksProto(bookmarkEntities)
	return bookmarks, err
}

func (u *bookmarkUsecase) Add(ctx context.Context, userID uint64, fqdn string) error {
	err := u.Repo.Create(ctx, userID, fqdn)
	return err
}

func (u *bookmarkUsecase) Delete(ctx context.Context, userID uint64, fqdn string) error {
	err := u.Repo.Delete(ctx, userID, fqdn)
	return err
}

func convToBookmarksProto(bookmarkEntities []*domain.Bookmark) ([]*pb.Bookmark, error) {
	var bookmarksProto []*pb.Bookmark
	for _, bookmarkEntity := range bookmarkEntities {
		pbCreatedAt, err := ptypes.TimestampProto(bookmarkEntity.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookmarkProto := &pb.Bookmark{
			UserId:    bookmarkEntity.UserID,
			Fqdn:      bookmarkEntity.FQDN,
			CreatedAt: pbCreatedAt,
		}
		bookmarksProto = append(bookmarksProto, bookmarkProto)
	}
	return bookmarksProto, nil
}
