package usecase

import (
	"context"
	"database/sql"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	"server/domain/model"
	"server/domain/repository"
	pb "server/etc/protocol"
	"server/infrastructure/store/mock"
	"testing"
)

func TestNewBookmarkUsecase(t *testing.T) {
	type args struct {
		bookmarkRepo repository.BookmarkRepository
	}
	tests := []struct {
		name string
		args args
		want BookmarkUsecase
	}{
		{
			"success",
			args{&mock.BookmarkStore{}},
			&bookmarkUsecase{&mock.BookmarkStore{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookmarkUsecase(tt.args.bookmarkRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookmarkUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookmarkUsecase_Add(t *testing.T) {
	type fields struct {
		Repo repository.BookmarkRepository
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
		{"success",
			fields{Repo: &mock.BookmarkStore{
				OnCreate: func(ctx context.Context, userID uint64, fqdn string) error {
					return nil
				},
			}},
			args{
				ctx:    context.Background(),
				userID: testdata_test.PbBookmark1.UserId,
				fqdn:   testdata_test.PbBookmark1.Fqdn,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &bookmarkUsecase{
				Repo: tt.fields.Repo,
			}
			if err := u.Add(tt.args.ctx, tt.args.userID, tt.args.fqdn); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookmarkUsecase_Delete(t *testing.T) {
	type fields struct {
		Repo repository.BookmarkRepository
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
		{"success",
			fields{Repo: &mock.BookmarkStore{
				OnDelete: func(ctx context.Context, userID uint64, fqdn string) error {
					return nil
				},
			}},
			args{
				ctx:    context.Background(),
				userID: testdata_test.PbBookmark1.UserId,
				fqdn:   testdata_test.PbBookmark1.Fqdn,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &bookmarkUsecase{
				Repo: tt.fields.Repo,
			}
			if err := u.Delete(tt.args.ctx, tt.args.userID, tt.args.fqdn); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookmarkUsecase_Get(t *testing.T) {
	type fields struct {
		Repo repository.BookmarkRepository
	}
	type args struct {
		ctx    context.Context
		userID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*pb.Stat
		wantErr bool
	}{
		{"bookmark exist",
			fields{Repo: &mock.BookmarkStore{
				OnGet: func(ctx context.Context, userID uint64) ([]*model.Stat, error) {
					return testdata_test.Stats, nil
				},
			}},
			args{
				ctx:    context.Background(),
				userID: testdata_test.PbUser1.Id,
			},
			testdata_test.PbStats,
			false,
		},
		{"bookmark not exist",
			fields{Repo: &mock.BookmarkStore{
				OnGet: func(ctx context.Context, userID uint64) ([]*model.Stat, error) {
					return nil, sql.ErrNoRows
				},
			}},
			args{
				ctx:    context.Background(),
				userID: testdata_test.PbUser1.Id,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &bookmarkUsecase{
				Repo: tt.fields.Repo,
			}
			got, err := u.Get(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
