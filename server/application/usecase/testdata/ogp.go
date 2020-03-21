package testdata

import (
	"server/domain/model"
	"time"
)

// entity
var Ogp1 = &model.Ogp{
	Date:    time.Date(2020, 10, 10, 10, 10, 10, 10, location),
	FQDN:    "example.com",
	Host:    "asia-n1",
	TweetID: 1,
	Type:    "website",
	Lang:    "ja",
}

// array
var Ogps = []*model.Ogp{Ogp1}
