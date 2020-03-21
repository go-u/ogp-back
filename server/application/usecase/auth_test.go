package usecase

import (
	"errors"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	"server/domain/service"
	"server/infrastructure/auth/mock"
	"testing"
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		authService service.AuthService
	}
	tests := []struct {
		name string
		args args
		want AuthUsecase
	}{
		{
			"success",
			args{&mock.AuthService{}},
			&authUsecase{&mock.AuthService{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.authService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_Verify(t *testing.T) {
	type fields struct {
		Service service.AuthService
	}
	type args struct {
		jwt string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"invalid jwt",
			fields{
				&mock.AuthService{
					OnVerify: func(jwt string) (string, error) {
						return "", errors.New("invalid jwt")
					},
				},
			},
			args{jwt: "invalid jwt"},
			"",
			true,
		},
		{
			"valid jwt",
			fields{
				&mock.AuthService{
					OnVerify: func(jwt string) (string, error) {
						return testdata_test.User1.UID, nil
					},
				},
			},
			args{jwt: "valid jwt"},
			testdata_test.User1.UID,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUsecase{
				Service: tt.fields.Service,
			}
			got, err := u.Verify(tt.args.jwt)
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
