package mock

import (
	"context"
	pb "server/etc/protocol"
)

type BookmarkUsecase struct {
	OnGet    func(context.Context, uint64) ([]*pb.Bookmark, error)
	OnAdd    func(context.Context, uint64, string) error
	OnDelete func(context.Context, uint64, string) error
}

func (u *BookmarkUsecase) Get(ctx context.Context, userID uint64) ([]*pb.Bookmark, error) {
	return u.OnGet(ctx, userID)
}

func (u *BookmarkUsecase) Add(ctx context.Context, userID uint64, fqdn string) error {
	return u.OnAdd(ctx, userID, fqdn)
}

func (u *BookmarkUsecase) Delete(ctx context.Context, userID uint64, fqdn string) error {
	return u.OnDelete(ctx, userID, fqdn)
}
