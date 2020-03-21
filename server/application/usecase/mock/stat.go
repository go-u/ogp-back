package mock

import (
	"context"
	pb "server/etc/protocol"
)

type StatUsecase struct {
	OnGet func(context.Context) ([]*pb.Stat, error)
}

func (u *StatUsecase) Get(ctx context.Context) ([]*pb.Stat, error) {
	return u.OnGet(ctx)
}
