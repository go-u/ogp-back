package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
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

func (s *BookmarkStore) Get(ctx context.Context, userID uint64) ([]*domain.Bookmark, error) {
	bookmarks, err := models.Bookmarks(models.BookmarkWhere.UserID.EQ(userID)).All(ctx, s.SqlHandler.Conn)
	var bookmarkEntities []*domain.Bookmark
	for _, bookmark := range bookmarks {
		bookmarkEntity := ConvToBookmarkEntity(bookmark)
		bookmarkEntities = append(bookmarkEntities, bookmarkEntity)
	}
	return bookmarkEntities, err
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
