package mock

import (
	"context"
	pb "server/etc/protocol"
)

type UserUsecase struct {
	OnGetByUID      func(context.Context, string) (*pb.User, error)
	OnCreateAccount func(context.Context, string, string) error
	OnDeleteAccount func(context.Context, string) error
}

func (u *UserUsecase) GetByUID(ctx context.Context, uid string) (*pb.User, error) {
	return u.OnGetByUID(ctx, uid)
}

func (u *UserUsecase) CreateAccount(ctx context.Context, uid string, name string) error {
	return u.OnCreateAccount(ctx, uid, name)
}

func (u *UserUsecase) DeleteAccount(ctx context.Context, uid string) error {
	return u.OnDeleteAccount(ctx, uid)
}
