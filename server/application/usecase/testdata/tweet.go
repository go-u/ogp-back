package testdata

import (
	"github.com/dghubble/go-twitter/twitter"
	pb "server/etc/protocol"
)

// oembed
var TweetEmbed1 = &twitter.OEmbedTweet{
	URL:          "example.com",
	ProviderURL:  "example.com",
	ProviderName: "example.com",
	AuthorName:   "author",
	Version:      "1",
	AuthorURL:    "example.com",
	Type:         "website",
	HTML:         "<p>html<p>",
	Height:       1000,
	Width:        1000,
	CacheAge:     "3600",
}

// pb
var PbTweet1 = &pb.Tweet{
	Url:  TweetEmbed1.URL,
	Html: TweetEmbed1.HTML,
}

// array
var TweetEmbeds = []*twitter.OEmbedTweet{TweetEmbed1}
var PbTweets = []*pb.Tweet{PbTweet1}
