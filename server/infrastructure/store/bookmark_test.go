package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"testing"
)

func TestNewBookmarkRepository(t *testing.T) {
	type args struct {
		sqlHandler SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.BookmarkRepository
	}{
		{
			"success",
			args{*NewSqlHandler("ogp-test")},
			&BookmarkStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookmarkRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewBookmarkRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookmarkStore_Create(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx    context.Context
		userID uint64
		fqdn   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"success",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:    context.Background(),
				userID: testdata_test.Bookmark1.UserID,
				fqdn:   testdata_test.Bookmark1.FQDN,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BookmarkStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.Create(tt.args.ctx, tt.args.userID, tt.args.fqdn); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBookmarkStore_Get(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx    context.Context
		userID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Stat
		wantErr bool
	}{
		{
			"success",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:    context.Background(),
				userID: testdata_test.User1.ID,
			},
			testdata_test.Stats,
			false,
		},
	}

	// insert
	newStat := models.Stat{
		Date:        testdata_test.Stat1.Date,
		FQDN:        testdata_test.Stat1.FQDN,
		Host:        testdata_test.Stat1.Host,
		Count:       testdata_test.Stat1.Count,
		Title:       testdata_test.Stat1.Title,
		Description: testdata_test.Stat1.Description,
		Image:       testdata_test.Stat1.Image,
		Type:        testdata_test.Stat1.Type,
		Lang:        testdata_test.Stat1.Lang,
	}

	sqlHandler := NewSqlHandler("ogp-test")
	err := newStat.Insert(context.Background(), sqlHandler.Conn, boil.Blacklist())
	if err != nil {
		log.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BookmarkStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.Get(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}

	// tear down
	deleteStat, err := models.Stats(models.StatWhere.FQDN.EQ(newStat.FQDN)).One(context.Background(), sqlHandler.Conn)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = deleteStat.Delete(context.Background(), sqlHandler.Conn)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestBookmarkStore_Delete(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx    context.Context
		userID uint64
		fqdn   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"success",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:    context.Background(),
				userID: testdata_test.Bookmark1.UserID,
				fqdn:   testdata_test.Bookmark1.FQDN,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BookmarkStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.Delete(tt.args.ctx, tt.args.userID, tt.args.fqdn); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
