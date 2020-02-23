package tools

import (
	"context"
	"database/sql"
	"net/http"
	"server/db/models"
	"server/initialize/db"
)

func GetUserFromJWT(r *http.Request) (*models.User, error) {
	uid, err := GetUidFromJWT(r)
	if err != nil {
		return nil, err
	}
	user, err := getUserFromUID(*uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUserFromUID(uid string) (*models.User, error) {
	user, err := models.Users(models.UserWhere.UID.EQ(uid)).One(context.Background(), db.Connection)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if user == nil || err == sql.ErrNoRows {
		return nil, nil
	}
	return user, nil
}
