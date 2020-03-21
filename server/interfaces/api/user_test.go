package api

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"server/application/usecase/mock"
	testdata_test "server/application/usecase/testdata"
	pb "server/etc/protocol"
	"strings"
	"testing"
)

func TestHandler_handleMe(t *testing.T) {
	type fields struct {
		Config *Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"valid jwt / valid user",
			fields{
				router: chi.NewRouter(),
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
			args{r: testdata_test.UserIdentifyMeRequest()},
			200,
			`{"id":1,"uid":"1","name":"user1","created_at":{"seconds":1602324610,"nanos":10}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}

func TestHandler_handleCreateAccount(t *testing.T) {
	type fields struct {
		Config *Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"valid jwt / valid name",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnCreateAccount: func(ctx context.Context, uid string, name string) error {
							return nil
						},
					},
				},
			},
			args{r: testdata_test.UserCreateRequest()},
			200,
			`"ok"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}

func TestHandler_handleDeleteAccount(t *testing.T) {
	type fields struct {
		Config *Config
		router chi.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			"valid jwt / valid user",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnDeleteAccount: func(ctx context.Context, uid string) error {
							return nil
						},
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return testdata_test.PbUser1, nil
						},
					},
				},
			},
			args{r: testdata_test.UserDeleteRequest()},
			200,
			`"ok"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.Config)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.args.r)

			if tt.wantCode != w.Code {
				t.Fatalf("test %q: want status code %d got %d", tt.name, tt.wantCode, w.Code)
			}

			if tt.wantBody != strings.TrimSpace(w.Body.String()) {
				t.Fatalf("test %q: want response body %q got %q", tt.name, tt.wantBody, w.Body)
			}
		})
	}
}
