package ogps

import (
	"context"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"net/http"
	"server/db/models"
	"server/initialize/db"
)

func Get(r *http.Request) ([]*models.Ogp, error) {
	return GetByPopular(r)

	// 高速に並び替え出来るようにフロントで処理したため以下は省略
	//order := r.URL.Query().Get("order")
	//if order == "popular" {
	//	return GetByPopular(r)
	//} else {
	//	return GetByNewer(r)
	//}
}

func GetByPopular(r *http.Request) ([]*models.Ogp, error) {
	ogps, err := models.Ogps(qm.OrderBy("-"+models.OgpColumns.Bookmarks)).All(context.Background(), db.Connection)
	return ogps, err
}

//func GetByNewer(r *http.Request) ([]*models.Ogp, error) {
//	ogps, err := models.Ogps(qm.OrderBy("-"+models.OgpColumns.CreatedAt)).All(context.Background(), db.Connection)
//	return ogps, err
//}
