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

func TestHandler_handleGetStats(t *testing.T) {
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
			"valid request",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					StatUsecase: &mock.StatUsecase{
						OnGet: func(context.Context) ([]*pb.Stat, error) {
							return testdata_test.PbStats, nil
						},
					},
				},
			},
			args{r: testdata_test.StatGetRequest()},
			200,
			`[{"fqdn":"example.com","count":30,"title":"title","description":"description","image":"https://example.com/image01.png","lang":"ja"}]`,
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
