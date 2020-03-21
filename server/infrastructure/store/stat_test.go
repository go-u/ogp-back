package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"reflect"
	"server/application/usecase/testdata"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
	"testing"
)

func TestNewStatRepository(t *testing.T) {
	type args struct {
		sqlHandler SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.StatRepository
	}{
		{
			"success",
			args{*NewSqlHandler("ogp-test")},
			&StatStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewStatRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatStore_Get_NotExist(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Stat
		wantErr bool
	}{
		{
			"no items",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx: context.Background(),
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.Get(tt.args.ctx)
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

func TestStatStore_Get_Exist(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Stat
		wantErr bool
	}{
		{
			"exist",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx: context.Background(),
			},
			[]*domain.Stat{testdata.Stat1},
			false,
		},
	}

	// insert
	newStat := models.Stat{
		Date:        testdata.Stat1.Date,
		FQDN:        testdata.Stat1.FQDN,
		Host:        testdata.Stat1.Host,
		Count:       testdata.Stat1.Count,
		Title:       testdata.Stat1.Title,
		Description: testdata.Stat1.Description,
		Image:       testdata.Stat1.Image,
		Type:        testdata.Stat1.Type,
		Lang:        testdata.Stat1.Lang,
	}

	sqlHandler := NewSqlHandler("ogp-test")
	err := newStat.Insert(context.Background(), sqlHandler.Conn, boil.Blacklist())
	if err != nil {
		log.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convToStatEntity(t *testing.T) {
	type args struct {
		stat *models.Stat
	}
	tests := []struct {
		name string
		args args
		want *domain.Stat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convToStatEntity(tt.args.stat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convToStatEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumStats(t *testing.T) {
	type args struct {
		stats []*models.Stat
	}
	tests := []struct {
		name string
		args args
		want []*models.Stat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumStats(tt.args.stats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
