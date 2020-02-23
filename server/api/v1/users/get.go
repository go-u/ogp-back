package users

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
)

func GetUserDetail(r *http.Request) (*models.User, error) {
	username := chi.URLParam(r, "username")
	user, err := models.Users(models.UserWhere.Username.EQ(username)).One(context.Background(), db.Connection)
	if user == nil || err != nil {
		return nil, err
	}
	return user, err
}

func GetSelfDetail(r *http.Request) (*models.User, error) {
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, err
}
