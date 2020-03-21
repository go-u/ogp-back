package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
)

func NewClient() *twitter.Client {
	return twitter.NewClient(http.DefaultClient)
}
