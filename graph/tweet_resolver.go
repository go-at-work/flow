package graph

import (
	"context"

	"github.com/arisromil/flow"
)

func mapTweet(t flow.Tweet) *Tweet {
	return &Tweet{
		ID:        t.ID,
		Body:      t.Body,
		UserID:    t.UserID,
		CreatedAt: t.CreatedAt,
	}
}

func mapTweets(tweets []flow.Tweet) []*Tweet {
	tt := make([]*Tweet, len(tweets))

	for i, t := range tweets {
		tt[i] = mapTweet(t)
	}
	return tt

}

func (q *queryResolver) Tweets(ctx context.Context) ([]*Tweet, error) {
	tweets, err := q.Resolver.TweetService.All(ctx)
	if err != nil {
		return nil, buildBadRequestError(ctx, err)
	}
	return mapTweets(tweets), nil

}

func (m *mutationResolver) CreateTweet(ctx context.Context, input flow.CreateTweetRequest) (*Tweet, error) {
	tweet, err := m.TweetService.CreateTweet(ctx, flow.CreateTweetRequest{
		Body: input.Body,
	})
	if err != nil {
		return nil, buildBadRequestError(ctx, err)
	}

	return mapTweet(tweet), nil
}
