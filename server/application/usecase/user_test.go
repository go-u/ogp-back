package usecase

import (
	"context"
	"database/sql"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	pb "server/etc/protocol"
	"server/infrastructure/store/mock"
	"testing"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		{
			"success",
			args{&mock.UserStore{}},
			&userUsecase{&mock.UserStore{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convToUserProto(t *testing.T) {
	type args struct {
		userEntity *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.User
		wantErr bool
	}{
		{
			"success",
			args{testdata_test.User1},
			testdata_test.PbUser1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convToUserProto(tt.args.userEntity)
			if (err != nil) != tt.wantErr {
				t.Errorf("convToUserProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convToUserProto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_CreateAccount(t *testing.T) {
	type fields struct {
		Repo repository.UserRepository
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
			fields{Repo: &mock.UserStore{
				OnCreate: func(ctx context.Context, uid string, name string) error {
					return nil
				},
			}},
			args{
				ctx:  context.Background(),
				uid:  testdata_test.PbUser1.Uid,
				name: testdata_test.PbUser1.Name,
			},
			false,
		},
		{
			"no display name",
			fields{Repo: &mock.UserStore{}},
			args{
				ctx:  context.Background(),
				uid:  testdata_test.PbUser1.Uid,
				name: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				Repo: tt.fields.Repo,
			}
			if err := u.CreateAccount(tt.args.ctx, tt.args.uid, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUsecase_DeleteAccount(t *testing.T) {
	type fields struct {
		Repo repository.UserRepository
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
			fields{Repo: &mock.UserStore{
				OnDelete: func(ctx context.Context, uid string) error {
					return nil
				},
			}},
			args{
				ctx: context.Background(),
				uid: testdata_test.PbUser1.Uid,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				Repo: tt.fields.Repo,
			}
			if err := u.DeleteAccount(tt.args.ctx, tt.args.uid); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUsecase_GetByUID(t *testing.T) {
	type fields struct {
		Repo repository.UserRepository
	}
	type args struct {
		ctx context.Context
		uid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		{
			"user exist",
			fields{Repo: &mock.UserStore{
				OnGetByUID: func(ctx context.Context, uid string) (*domain.User, error) {
					return testdata_test.User1, nil
				},
			}},
			args{
				ctx: context.Background(),
				uid: "1",
			},
			testdata_test.PbUser1,
			false,
		},
		{
			"user not exist",
			fields{Repo: &mock.UserStore{
				OnGetByUID: func(ctx context.Context, uid string) (*domain.User, error) {
					return nil, sql.ErrNoRows
				},
			}},
			args{
				ctx: context.Background(),
				uid: "2",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				Repo: tt.fields.Repo,
			}
			got, err := u.GetByUID(tt.args.ctx, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidDisplayName(t *testing.T) {
	type args struct {
		displayname string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"blank name",
			args{displayname: ""},
			false,
		},
		{
			"minimum length: 1",
			args{displayname: "I"},
			true,
		},
		{
			"length 4",
			args{displayname: "name"},
			true,
		},
		{
			"max length: 50",
			args{displayname: "01234567890123456789012345678901234567890123456789"},
			true,
		},
		{
			"over limit: 51",
			args{displayname: "01234567890123456789012345678901234567890123456789A"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDisplayName(tt.args.displayname); got != tt.want {
				t.Errorf("isValidDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}
