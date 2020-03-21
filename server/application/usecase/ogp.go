package usecase

import (
	"context"
	"server/domain/repository"
	"server/domain/service"
	pb "server/etc/protocol"
)

type OgpUsecase interface {
	Get(context.Context, string) ([]*pb.Tweet, error)
}

type ogpUsecase struct {
	Repo    repository.OgpRepository
	Service service.TwitterService
}

func NewOgpUsecase(ogpRepo repository.OgpRepository, tweetService service.TwitterService) OgpUsecase {
	ogpUsecase := ogpUsecase{
		ogpRepo,
		tweetService,
	}
	return &ogpUsecase
}

func (u *ogpUsecase) Get(ctx context.Context, fqdn string) ([]*pb.Tweet, error) {
	ogpEntities, err := u.Repo.Get(ctx, fqdn)
	if err != nil {
		return make([]*pb.Tweet, 0), err
	}

	var tweetIds []int64
	for _, ogpEntity := range ogpEntities {
		tweetIds = append(tweetIds, ogpEntity.TweetID)
	}

	oEmbeds, err := u.Service.GetTweets(tweetIds)
	if err != nil {
		return make([]*pb.Tweet, 0), err
	}

	var tweets []*pb.Tweet
	for _, oEmbed := range oEmbeds {
		tweet := &pb.Tweet{
			Url:  oEmbed.URL,
			Html: oEmbed.HTML,
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
