package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"server/domain/service"
)

type TwitterService struct {
	Client twitter.Client
}

func NewTwitterService(client twitter.Client) service.TwitterService {
	twitterService := TwitterService{client}
	return &twitterService
}

func (s *TwitterService) GetTweets(ids []int64) ([]*twitter.OEmbedTweet, error) {
	var oEmbeds = []*twitter.OEmbedTweet{}
	for _, id := range ids {
		True := true
		False := false
		params := twitter.StatusOEmbedParams{
			ID:         id,
			MaxWidth:   0,
			HideMedia:  &False,
			HideThread: &True,
			OmitScript: &True,
			WidgetType: "video",
			HideTweet:  nil,
		}
		emb, _, err := s.Client.Statuses.OEmbed(&params)
		if err != nil {
			log.Println(err) // 該当ツイートが削除されていた場合 422:twitter: 34 Sorry, that page does not exist
			continue
			//return oEmbeds, err
		}
		oEmbeds = append(oEmbeds, emb)
	}
	return oEmbeds, nil
}
