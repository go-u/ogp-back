package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/api/v1/media"
	"server/config"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
)

func Register(r *http.Request) (*models.User, error) {
	register_info := RegisterInfo{}
	err := render.DecodeJSON(r.Body, &register_info)
	if err != nil {
		return nil, errors.New("Binding Failed")
	}
	// user already exist
	user_in_db, err := tools.GetUserFromJWT(r)
	if err != nil && err != sql.ErrNoRows {
		return nil, err // pure error
	}
	if user_in_db != nil {
		return nil, errors.New("user already exist")
	}
	// valid jwt token
	uid, err := tools.GetUidFromJWT(r)
	if err != nil {
		return nil, err
	}
	provider, err := tools.GetSiginProviderFromJWT(r)
	if err != nil {
		return nil, err
	}

	new_user := &models.User{
		UID:         *uid,
		Username:    tools.GenerateRandomAsciiString(10),
		Displayname: register_info.Displayname,
		Provider:    *provider,
	}

	// decode base64 to bytes
	image_bytes, err := media.DecodeBase64ImageToBytes(register_info.AvatarBase64)
	if err != nil {
		return nil, err
	}

	tx, err := db.Connection.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	err = new_user.Insert(context.Background(), db.Connection, boil.Infer())
	if err != nil {
		_ = tx.Rollback() // rollback
		return nil, err
	}

	err = UploadAvatar(*new_user, image_bytes, "avatar", config.AvatarWidth, config.AvatarHeight)
	if err != nil {
		_ = tx.Rollback() // rollback
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	//user_detail_musked := CovertUser_To_MuskedUser(new_user)
	return new_user, nil
}
