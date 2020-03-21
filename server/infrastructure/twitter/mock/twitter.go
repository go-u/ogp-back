package mock

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterService struct {
	OnGetTweets func(ids []int64) ([]*twitter.OEmbedTweet, error)
}

func (s *TwitterService) GetTweets(ids []int64) ([]*twitter.OEmbedTweet, error) {
	return s.OnGetTweets(ids)
}
