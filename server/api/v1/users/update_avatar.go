package users

import (
	"context"
	"fmt"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/api/v1/media"
	"server/config"
	"server/db/models"
	"server/initialize/db"
	"server/initialize/gcs"
	"server/tools"
	"time"
)

func UpdateAvatar(r *http.Request) error {
	image_byte, err := getByteImageFromRequest(r)
	if err != nil {
		return err
	}
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return err
	}
	return UploadAvatar(*user, image_byte, "avatar", config.AvatarWidth, config.AvatarHeight)
}

func getByteImageFromRequest(r *http.Request) ([]byte, error) {
	// binding
	base64_image := media.Base64Image{}
	err := render.DecodeJSON(r.Body, &base64_image)
	if err != nil {
		return nil, err
	}
	// decode base64 to bytes
	image_bytes, err := media.DecodeBase64ImageToBytes(base64_image.Image)
	if err != nil {
		return nil, err
	}
	return image_bytes, nil
}

func UploadAvatar(user models.User, image_bytes []byte, item_type string, width uint, height uint) error {
	// https://storage.googleapis.com/boost-prd-avatar/user/3/1576899090.png
	path_without_extension := fmt.Sprintf("user/%d/%d", user.ID, time.Now().Unix())
	Bucket := gcs.Buckets.Avatar

	// delete old image with all of dir
	delete_dir_path := fmt.Sprintf("user/%d/", user.ID)
	err := media.DeleteDirGCS(Bucket, delete_dir_path)
	if err != nil {
		return err
	}

	// upload
	media_link, err := media.UploadImageGCS(Bucket, image_bytes, path_without_extension, true, width, height, 60*60*24)
	if err != nil {
		return err
	}

	// update db
	user.Avatar = *media_link
	_, err = user.Update(context.Background(), db.Connection, boil.Whitelist(models.UserColumns.Avatar, models.UserColumns.UpdatedAt))
	if err != nil {
		return err
	}
	return nil
}
