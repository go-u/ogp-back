package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	"server/domain/service"
	pb "server/etc/protocol"
	"server/infrastructure/store/mock"
	twimock "server/infrastructure/twitter/mock"
	"testing"
)

func TestNewOgpUsecase(t *testing.T) {
	type args struct {
		ogpRepo      repository.OgpRepository
		tweetService service.TwitterService
	}
	tests := []struct {
		name string
		args args
		want OgpUsecase
	}{
		{
			"success",
			args{
				&mock.OgpStore{},
				&twimock.TwitterService{},
			},
			&ogpUsecase{
				&mock.OgpStore{},
				&twimock.TwitterService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOgpUsecase(tt.args.ogpRepo, tt.args.tweetService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOgpUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ogpUsecase_Get(t *testing.T) {
	type fields struct {
		Repo    repository.OgpRepository
		Service service.TwitterService
	}
	type args struct {
		ctx  context.Context
		fqdn string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*pb.Tweet
		wantErr bool
	}{
		{"ogp exist",
			fields{
				Repo: &mock.OgpStore{
					OnGet: func(ctx context.Context, fqdn string) ([]*domain.Ogp, error) {
						return testdata_test.Ogps, nil
					},
				},
				Service: &twimock.TwitterService{
					OnGetTweets: func(ids []int64) ([]*twitter.OEmbedTweet, error) {
						return testdata_test.TweetEmbeds, nil
					},
				},
			},
			args{
				ctx: context.Background(),
			},
			testdata_test.PbTweets,
			false,
		},
		{"ogp not exist",
			fields{
				Repo: &mock.OgpStore{
					OnGet: func(ctx context.Context, fqdn string) ([]*domain.Ogp, error) {
						return nil, sql.ErrNoRows
					},
				},
				Service: &twimock.TwitterService{
					OnGetTweets: func(ids []int64) ([]*twitter.OEmbedTweet, error) {
						return nil, errors.New("api error")
					},
				},
			},
			args{
				ctx: context.Background(),
			},
			make([]*pb.Tweet, 0),
			true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ogpUsecase{
				Repo:    tt.fields.Repo,
				Service: tt.fields.Service,
			}
			got, err := u.Get(tt.args.ctx, tt.args.fqdn)
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
