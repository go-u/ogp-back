package usecase

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	domain "server/domain/model"
	"server/domain/repository"
	"server/etc/config"
	pb "server/etc/protocol"
)

type UserUsecase interface {
	GetByUID(context.Context, string) (*pb.User, error)
	CreateAccount(context.Context, string, string) error
	DeleteAccount(context.Context, string) error
}

type userUsecase struct {
	Repo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo}
	return &userUsecase
}

func (u *userUsecase) GetByUID(ctx context.Context, uid string) (*pb.User, error) {
	userEntity, err := u.Repo.GetByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return convToUserProto(userEntity)
}

func (u *userUsecase) CreateAccount(ctx context.Context, uid string, name string) error {
	isValid := isValidDisplayName(name)
	if !isValid {
		return errors.New("invalid display name")
	}

	return u.Repo.Create(ctx, uid, name)
}

func (u *userUsecase) DeleteAccount(ctx context.Context, uid string) error {
	return u.Repo.Delete(ctx, uid)
}

func convToUserProto(userEntity *domain.User) (*pb.User, error) {
	cretedAt, err := ptypes.TimestampProto(userEntity.CreatedAt)
	if err != nil {
		return nil, err
	}
	pbUser := &pb.User{
		Id:        userEntity.ID,
		Uid:       userEntity.UID,
		Name:      userEntity.Name,
		CreatedAt: cretedAt,
	}
	return pbUser, nil
}

func isValidDisplayName(displayname string) bool {
	return config.DisplayNameRegexp.MatchString(displayname)
}
