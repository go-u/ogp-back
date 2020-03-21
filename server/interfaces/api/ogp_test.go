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

func TestHandler_handleGetOgps(t *testing.T) {
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
			"valid fqdn",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					OgpUsecase: &mock.OgpUsecase{
						OnGet: func(context.Context, string) ([]*pb.Tweet, error) {
							return testdata_test.PbTweets, nil
						},
					},
				},
			},
			args{r: testdata_test.OgpGetRequest()},
			200,
			`[{"url":"example.com","html":"\u003cp\u003ehtml\u003cp\u003e"}]`,
		},
		{
			"invalid fqdn",
			fields{
				router: chi.NewRouter(),
				Config: &Config{
					OgpUsecase: &mock.OgpUsecase{
						OnGet: func(context.Context, string) ([]*pb.Tweet, error) {
							return nil, errors.New("wrong fqdn")
						},
					},
				},
			},
			args{r: testdata_test.OgpGetRequest()},
			422,
			`{"status":"Error rendering response.","error":"wrong fqdn"}`,
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
