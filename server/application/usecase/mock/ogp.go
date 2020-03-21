package mock

import (
	"context"
	pb "server/etc/protocol"
)

type OgpUsecase struct {
	OnGet func(context.Context, string) ([]*pb.Tweet, error)
}

func (u *OgpUsecase) Get(ctx context.Context, fqdn string) ([]*pb.Tweet, error) {
	return u.OnGet(ctx, fqdn)
}
