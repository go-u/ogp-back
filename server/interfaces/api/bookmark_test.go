package api

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"server/application/usecase/mock"
	testdata_test "server/application/usecase/testdata"
	pb "server/etc/protocol"
	"strings"
	"testing"
)

func TestHandler_handleAddBookmark(t *testing.T) {
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
			"valid jwt / valid bookmark",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					BookmarkUsecase: &mock.BookmarkUsecase{
						OnAdd: func(context.Context, uint64, string) error {
							return nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return testdata_test.PbUser1, nil
						},
					},
				},
			},
			args{r: testdata_test.BookmarkAddRequest()},
			200,
			`"ok"`,
		},
		{
			"wrong jwt / valid bookmark",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return "", errors.New("invalid jwt")
						},
					},
				},
			},
			args{r: testdata_test.BookmarkAddRequest()},
			401,
			`{"status":"Unauthorized","error":"Authentication required, invalid jwt"}`,
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

func TestHandler_handleDeleteBookmark(t *testing.T) {
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
			"valid jwt / valid bookmark",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					BookmarkUsecase: &mock.BookmarkUsecase{
						OnDelete: func(context.Context, uint64, string) error {
							return nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return testdata_test.PbUser1, nil
						},
					},
				},
			},
			args{r: testdata_test.BookmarkDeleteRequest()},
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

func TestHandler_handleGetBookmarks(t *testing.T) {
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
			"valid jwt / valid bookmark",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					AuthUsecase: &mock.AuthUsecase{
						OnVerify: func(jwt string) (string, error) {
							return testdata_test.User1.UID, nil
						},
					},
					BookmarkUsecase: &mock.BookmarkUsecase{
						OnGet: func(context.Context, uint64) ([]*pb.Bookmark, error) {
							return testdata_test.PbBookmarks, nil
						},
					},
					UserUsecase: &mock.UserUsecase{
						OnGetByUID: func(ctx context.Context, uid string) (*pb.User, error) {
							return testdata_test.PbUser1, nil
						},
					},
				},
			},
			args{r: testdata_test.BookmarkGetRequest()},
			200,
			`[{"user_id":1,"fqdn":"example.com","created_at":{"seconds":1602324610,"nanos":10}}]`,
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
