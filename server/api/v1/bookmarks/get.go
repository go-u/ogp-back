package bookmarks

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/queries"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
)

func GetBookmarks(r *http.Request) ([]*models.Ogp, error) {
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}

	bookmarks, err := models.Bookmarks(models.BookmarkWhere.UserID.EQ(user.ID)).All(context.Background(), db.Connection)
	if bookmarks == nil || err != nil {
		return make([]*models.Ogp, 0), err
	}

	query_str_before := `SELECT * FROM ogp WHERE `
	query_str_middle := ``
	query_str_after := `;`

	for _, bookmark := range bookmarks {
		query_str_middle += fmt.Sprintf("(id = %d) OR ", bookmark.OgpID)
	}
	query_str_middle = query_str_middle[:(len(query_str_middle) - 3)]
	query_str := query_str_before + query_str_middle + query_str_after
	ogps := []*models.Ogp{}
	err = queries.Raw(query_str).Bind(context.Background(), db.Connection, &ogps)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return ogps, nil
}
