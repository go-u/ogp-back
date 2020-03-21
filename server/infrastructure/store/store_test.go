package store

import (
	"log"
	"os"
	"reflect"
	"server/etc/utils"
	"testing"
)

func TestMain(m *testing.M) {
	println("setup...")
	// create table
	sqlHandler := NewSqlHandler("ogp-test")
	err := utils.ResetAndMigrateDB(sqlHandler.Conn)
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	println("tear down...")

	os.Exit(code)
}

func TestNewSqlHandler(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name string
		args args
		want *SqlHandler
	}{
		{
			"success",
			args{projectID: "ogp-test"},
			&SqlHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSqlHandler(tt.args.projectID); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewSqlHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
