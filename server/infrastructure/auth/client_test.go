package auth

import (
	fbauth "firebase.google.com/go/auth"
	"reflect"
	secrets_firebase "server/etc/secrets/firebase"
	"server/etc/utils"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		PROJECT_ID string
	}
	tests := []struct {
		name string
		args args
		want *fbauth.Client
	}{
		{
			"local",
			args{PROJECT_ID: "ogp-local"},
			&fbauth.Client{},
		},
		{
			"stg",
			args{PROJECT_ID: "ogp-stg"},
			&fbauth.Client{},
		},
		{
			"prd",
			args{PROJECT_ID: "ogp-prd"},
			&fbauth.Client{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.PROJECT_ID); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createClient(t *testing.T) {
	type args struct {
		CRED_PATH string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"local",
			args{CRED_PATH: utils.GetMainPath() + secrets_firebase.FireBaseLocal},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createClient(tt.args.CRED_PATH)
			if (err != nil) != tt.wantErr {
				t.Errorf("createClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got != nil, tt.want) {
				t.Errorf("createClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCred(t *testing.T) {
	type args struct {
		PROJECT_ID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"local",
			args{PROJECT_ID: "ogp-local"},
			utils.GetMainPath() + secrets_firebase.FireBaseLocal,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCred(tt.args.PROJECT_ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCred() got = %v, want %v", got, tt.want)
			}
		})
	}
}
