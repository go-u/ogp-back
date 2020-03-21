package auth

import (
	"firebase.google.com/go/auth"
	"reflect"
	"server/domain/service"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		client auth.Client
	}
	tests := []struct {
		name string
		args args
		want service.AuthService
	}{
		{
			"success",
			args{client: auth.Client{}},
			&authService{auth.Client{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_Verify(t *testing.T) {
	type fields struct {
		Client auth.Client
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"valid",
			fields{*NewClient("ogp-local")},
			args{token: "token1"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authService{
				Client: tt.fields.Client,
			}
			got, err := a.Verify(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}
