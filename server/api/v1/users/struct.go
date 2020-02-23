package users

import (
	"server/db/models"
)

type RegisterInfo struct {
	Displayname  string `boil:"displayname" json:"displayname" toml:"displayname" yaml:"displayname"`
	AvatarBase64 string `boil:"avatar_base64" json:"avatar_base64" toml:"avatar_base64" yaml:"avatar_base64"`
}

type UpdateRequest struct {
	User         models.User `json:"user"`
	AvatarBase64 string      `json:"avatar_base64"`
	HeaderBase64 string      `json:"header_base64"`
}
