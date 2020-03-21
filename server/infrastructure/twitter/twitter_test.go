package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"reflect"
	"server/domain/service"
	"testing"
)

func TestNewTwitterService(t *testing.T) {
	type args struct {
		client twitter.Client
	}
	tests := []struct {
		name string
		args args
		want service.TwitterService
	}{
		{
			"success",
			args{*NewClient()},
			&TwitterService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTwitterService(tt.args.client); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewTwitterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwitterService_GetTweets(t *testing.T) {
	type fields struct {
		Client twitter.Client
	}
	type args struct {
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			"no ids",
			fields{*NewClient()},
			args{ids: nil},
			0,
			false,
		},
		{
			"wrong ids",
			fields{*NewClient()},
			args{ids: []int64{1}},
			0,
			false,
		},
		{
			"correct ids",
			fields{*NewClient()},
			args{ids: []int64{1232636643821182976}},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TwitterService{
				Client: tt.fields.Client,
			}
			got, err := s.GetTweets(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTweets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetTweets() got = %v, want %v", got, tt.want)
			}
		})
	}
}
