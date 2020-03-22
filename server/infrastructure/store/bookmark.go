package store

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"time"
)

func NewBookmarkRepository(sqlHandler SqlHandler) repository.BookmarkRepository {
	bookmarkRepository := BookmarkStore{sqlHandler}
	return &bookmarkRepository
}

type BookmarkStore struct {
	SqlHandler
}

func (s *BookmarkStore) Get(ctx context.Context, userID uint64) ([]*domain.Stat, error) {
	bookmarks, err := models.Bookmarks(models.BookmarkWhere.UserID.EQ(userID)).All(ctx, s.SqlHandler.Conn)
	var statEntities []*domain.Stat

	// todo: 以下のForで回す処理は N+1 で遅くなるため生SQLクエリ一発で取得するように変更(現状は件数が少ないので問題ない)
	for _, bookmark := range bookmarks {
		recent := time.Now().AddDate(0, 0, -2)
		stat, err := models.Stats(models.StatWhere.FQDN.EQ(bookmark.FQDN), models.StatWhere.Date.GTE(recent), qm.OrderBy("-"+models.StatColumns.Date), qm.Limit(1)).One(context.Background(), s.SqlHandler.Conn)
		if err == sql.ErrNoRows {
			continue
		} else if err != nil {
			return make([]*domain.Stat, 0), err
		}
		statEntity := convToStatEntity(stat)
		statEntities = append(statEntities, statEntity)
	}
	return statEntities, err
}

func (s *BookmarkStore) Create(ctx context.Context, userID uint64, fqdn string) error {
	newBookmark := models.Bookmark{
		UserID:    userID,
		FQDN:      fqdn,
		CreatedAt: time.Now(),
	}
	err := newBookmark.Insert(ctx, s.SqlHandler.Conn, boil.Blacklist())
	return err
}

func (s *BookmarkStore) Delete(ctx context.Context, userID uint64, fqdn string) error {
	bookmark, err := models.Bookmarks(models.BookmarkWhere.UserID.EQ(userID), models.BookmarkWhere.FQDN.EQ(fqdn)).One(ctx, s.SqlHandler.Conn)
	if err != nil {
		return err
	}
	_, err = bookmark.Delete(ctx, s.SqlHandler.Conn)
	return err
}

func ConvToBookmarkEntity(bookmark *models.Bookmark) *domain.Bookmark {
	bookmarkEntity := &domain.Bookmark{
		UserID:    bookmark.UserID,
		FQDN:      bookmark.FQDN,
		CreatedAt: bookmark.CreatedAt,
	}
	return bookmarkEntity
}
