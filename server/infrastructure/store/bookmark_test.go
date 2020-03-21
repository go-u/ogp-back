package store

import (
	"context"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
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
		want    []*domain.Bookmark
		wantErr bool
	}{
		{
			"success",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:    context.Background(),
				userID: testdata_test.User1.ID,
			},
			testdata_test.Bookmarks,
			false,
		},
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
