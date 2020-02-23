package users

import (
	"context"
	"errors"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
	"time"
)

func Update(r *http.Request) (*models.User, error) {
	update_request := models.User{}
	err := render.DecodeJSON(r.Body, &update_request)
	if err != nil {
		return nil, errors.New("Binding Failed")
	}

	// restrict user
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}
	if update_request.ID != user.ID {
		return nil, errors.New("own self only")
	}
	update_request.ID = user.ID

	// validate
	valid := isValidDisplayName(update_request.Displayname)
	if !valid {
		return nil, errors.New("wrong displayname")
	}

	// validate username if chaneged
	if user.Username != update_request.Username {
		valid = isValidUsername(update_request.Username)
		if !valid {
			return nil, errors.New("wrong username")
		}
	}

	// update
	update_request.UpdatedAt = time.Now()
	_, err = update_request.Update(context.Background(), db.Connection, boil.Whitelist(
		models.UserColumns.Username,
		models.UserColumns.Displayname,
		models.UserColumns.UpdatedAt,
	))
	if err != nil {
		return nil, err
	}
	return user, nil
}
