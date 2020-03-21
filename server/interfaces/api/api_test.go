package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"reflect"
	"server/application/usecase/mock"
	testdata_test "server/application/usecase/testdata"
	pb "server/etc/protocol"
	"testing"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			"success",
			args{
				&Config{
					AuthUsecase:     &mock.AuthUsecase{},
					BookmarkUsecase: &mock.BookmarkUsecase{},
					OgpUsecase:      &mock.OgpUsecase{},
					StatUsecase:     &mock.StatUsecase{},
					UserUsecase:     &mock.UserUsecase{},
				},
			},
			&Handler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.config); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_requestUser(t *testing.T) {
	type fields struct {
		Config *Config
		router chi.Router
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		{
			"valid jwt",
			fields{
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return testdata_test.PbUser1, nil
						},
					},
				},
			},
			args{r: testdata_test.BlankRequestWithSampleJwt()},
			testdata_test.PbUser1,
			false,
		},
		{
			"invalid jwt",
			fields{
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return "", errors.New("invalid jwt")
						},
					},
				},
			},
			args{r: testdata_test.BlankRequestWithSampleJwt()},
			nil,
			true,
		},
		{
			"valid jwt / no user exist in store",
			fields{
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return nil, sql.ErrNoRows
						},
					},
				},
			},
			args{r: testdata_test.BlankRequestWithSampleJwt()},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Config: tt.fields.Config,
				router: tt.fields.router,
			}
			got, err := h.requestUser(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
