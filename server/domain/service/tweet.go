package service

import "github.com/dghubble/go-twitter/twitter"

type TwitterService interface {
	GetTweets(ids []int64) ([]*twitter.OEmbedTweet, error)
}
