package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"time"
)

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserStore{sqlHandler}
	return &userRepository
}

type UserStore struct {
	SqlHandler
}

func (s *UserStore) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	repositoryUser, err := models.Users(models.UserWhere.UID.EQ(uid)).One(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:        repositoryUser.ID,
		UID:       repositoryUser.UID,
		Name:      repositoryUser.Name,
		CreatedAt: time.Time{},
	}
	return user, err
}

func (s *UserStore) Create(ctx context.Context, uid string, name string) error {
	user := models.User{
		UID:  uid,
		Name: name,
	}
	err := user.Insert(ctx, s.SqlHandler.Conn, boil.Infer())
	return err
}

func (s *UserStore) Delete(ctx context.Context, uid string) error {
	user, err := models.Users(models.UserWhere.UID.EQ(uid)).One(ctx, s.SqlHandler.Conn)
	if err != nil {
		return err
	}
	_, err = user.Delete(ctx, s.SqlHandler.Conn)
	return err
}
