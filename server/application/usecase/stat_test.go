package usecase

import (
	"context"
	"database/sql"
	"reflect"
	testdata_test "server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	pb "server/etc/protocol"
	"server/infrastructure/store/mock"
	"testing"
)

func TestNewStatUsecase(t *testing.T) {
	type args struct {
		statRepo repository.StatRepository
	}
	tests := []struct {
		name string
		args args
		want StatUsecase
	}{
		{
			"success",
			args{&mock.StatStore{}},
			&statUsecase{&mock.StatStore{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatUsecase(tt.args.statRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convToStatsProto(t *testing.T) {
	type args struct {
		statEntities []*domain.Stat
	}
	tests := []struct {
		name string
		args args
		want []*pb.Stat
	}{
		{
			"success",
			args{testdata_test.Stats},
			testdata_test.PbStats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convToStatsProto(tt.args.statEntities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convToStatsProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statUsecase_Get(t *testing.T) {
	type fields struct {
		Repo repository.StatRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*pb.Stat
		wantErr bool
	}{
		{"stat exist",
			fields{Repo: &mock.StatStore{
				OnGet: func(ctx context.Context) ([]*domain.Stat, error) {
					return testdata_test.Stats, nil
				},
			}},
			args{
				ctx: context.Background(),
			},
			testdata_test.PbStats,
			false,
		},
		{"stat not exist",
			fields{Repo: &mock.StatStore{
				OnGet: func(ctx context.Context) ([]*domain.Stat, error) {
					return nil, sql.ErrNoRows
				},
			}},
			args{
				ctx: context.Background(),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &statUsecase{
				Repo: tt.fields.Repo,
			}
			got, err := u.Get(tt.args.ctx)
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
