package ogps

import (
	"context"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
)

func Delete(r *http.Request) error {
	ogp := models.Ogp{}
	err := render.DecodeJSON(r.Body, &ogp)
	if err != nil {
		return err
	}

	// confirm ie exist
	ogp_in_db, err := models.FindOgp(context.Background(), db.Connection, ogp.ID)
	if err != nil {
		return err
	}

	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return err
	}

	if ogp_in_db.UserID != user.ID {
		return errors.New("not owner")
	}

	rowAff, err := ogp_in_db.Delete(context.Background(), db.Connection)
	if err != nil || rowAff == 0 {
		return err
	}

	return nil
}
