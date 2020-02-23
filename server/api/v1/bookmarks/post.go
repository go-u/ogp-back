package bookmarks

import (
	"context"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"server/db/models"
	"server/initialize/db"
	"server/tools"
	"time"
)

func Add(r *http.Request) (*models.Bookmark, error) {
	ogp := models.Ogp{}
	err := render.DecodeJSON(r.Body, &ogp)
	if err != nil {
		return nil, err
	}
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return nil, err
	}

	bookmark := models.Bookmark{
		UserID:    user.ID,
		OgpID:     ogp.ID,
		CreatedAt: time.Now(),
	}

	// confirm ogp exist
	ogp_in_db, err := models.FindOgp(context.Background(), db.Connection, ogp.ID)
	if err != nil {
		return nil, err
	}

	tx, err := db.Connection.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	err = bookmark.Insert(context.Background(), tx, boil.Blacklist())
	if err != nil {
		_ = tx.Rollback() // rollback
		return nil, err
	}
	// count up
	// クエリを使ってDB上でインクリメントした方が良いが、今回はプロトタイプのため簡略化
	ogp_in_db.Bookmarks += 1
	_, err = ogp_in_db.Update(context.Background(), tx, boil.Whitelist(models.OgpColumns.Bookmarks, models.OgpColumns.UpdatedAt))
	if err != nil {
		_ = tx.Rollback() // rollback
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &bookmark, err

}

func Delete(r *http.Request) error {
	ogp := models.Ogp{}
	err := render.DecodeJSON(r.Body, &ogp)
	if err != nil {
		return err
	}
	user, err := tools.GetUserFromJWT(r)
	if err != nil {
		return err
	}

	// confirm ogp exist
	ogp_in_db, err := models.FindOgp(context.Background(), db.Connection, ogp.ID)
	if err != nil {
		return err
	}

	bookmark, err := models.Bookmarks(models.BookmarkWhere.UserID.EQ(user.ID), models.BookmarkWhere.OgpID.EQ(ogp.ID)).One(context.Background(), db.Connection)
	if err != nil {
		return err
	}

	tx, err := db.Connection.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	_, err = bookmark.Delete(context.Background(), tx)
	if err != nil {
		_ = tx.Rollback() // rollback
		return err
	}

	// クエリを使ってDB上でデクリメントした方が良いが、今回はプロトタイプのため簡略化
	ogp_in_db.Bookmarks -= 1
	_, err = ogp_in_db.Update(context.Background(), tx, boil.Whitelist(models.OgpColumns.Bookmarks, models.OgpColumns.UpdatedAt))
	if err != nil {
		_ = tx.Rollback() // rollback
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
