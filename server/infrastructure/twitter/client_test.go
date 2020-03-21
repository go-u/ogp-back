package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *twitter.Client
	}{
		{
			"success",
			&twitter.Client{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
