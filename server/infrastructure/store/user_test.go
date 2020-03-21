package store

import (
	"context"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		sqlHandler SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.UserRepository
	}{
		{
			"success",
			args{*NewSqlHandler("ogp-test")},
			&UserStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStore_Create(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx  context.Context
		uid  string
		name string
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
				ctx:  context.Background(),
				uid:  testdata_test.User1.UID,
				name: testdata_test.User1.Name,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.Create(tt.args.ctx, tt.args.uid, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserStore_GetByUID(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx context.Context
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			"success",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx: context.Background(),
				uid: testdata_test.User1.UID,
			},
			testdata_test.User1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.GetByUID(tt.args.ctx, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || !reflect.DeepEqual(got.UID, tt.want.UID) {
				t.Errorf("GetByUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStore_Delete(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx context.Context
		uid string
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
				ctx: context.Background(),
				uid: testdata_test.User1.UID,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			if err := s.Delete(tt.args.ctx, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
