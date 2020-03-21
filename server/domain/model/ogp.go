package model

import "time"

type Ogp struct {
	Date    time.Time
	FQDN    string
	Host    string
	TweetID int64
	Type    string
	Lang    string
}
