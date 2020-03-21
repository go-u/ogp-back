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

func TestNewOgpRepository(t *testing.T) {
	type args struct {
		sqlHandler SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.OgpRepository
	}{
		{
			"success",
			args{*NewSqlHandler("ogp-test")},
			&OgpStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOgpRepository(tt.args.sqlHandler); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewOgpRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOgpStore_Get_NotExist(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx  context.Context
		fqdn string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Ogp
		wantErr bool
	}{
		{
			"no items",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:  context.Background(),
				fqdn: "",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OgpStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.Get(tt.args.ctx, tt.args.fqdn)
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

func TestOgpStore_Get_Exist(t *testing.T) {
	type fields struct {
		SqlHandler SqlHandler
	}
	type args struct {
		ctx  context.Context
		fqdn string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Ogp
		wantErr bool
	}{
		{
			"exist",
			fields{SqlHandler: *NewSqlHandler("ogp-test")},
			args{
				ctx:  context.Background(),
				fqdn: testdata.Ogp1.FQDN,
			},
			[]*domain.Ogp{testdata.Ogp1},
			false,
		},
	}

	// insert
	newOgp := models.Ogp{
		Date:    testdata.Ogp1.Date,
		FQDN:    testdata.Ogp1.FQDN,
		Host:    testdata.Ogp1.Host,
		TweetID: testdata.Ogp1.TweetID,
		Type:    testdata.Ogp1.Type,
		Lang:    testdata.Ogp1.Lang,
	}
	sqlHandler := NewSqlHandler("ogp-test")
	err := newOgp.Insert(context.Background(), sqlHandler.Conn, boil.Blacklist())
	if err != nil {
		log.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OgpStore{
				SqlHandler: tt.fields.SqlHandler,
			}
			got, err := s.Get(tt.args.ctx, tt.args.fqdn)
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
