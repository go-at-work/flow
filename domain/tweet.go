package domain

import (
	"context"

	"github.com/arisromil/flow"
)

type TweetService struct {
	TweetRepository flow.TweetRepository
}

func NewTweetService(repo flow.TweetRepository) *TweetService {
	return &TweetService{
		TweetRepository: repo,
	}
}

func (s *TweetService) All(ctx context.Context) ([]*flow.Tweet, error) {
	tweets, err := s.TweetRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (s *TweetService) CreateTweet(ctx context.Context, input flow.CreateTweetRequest) (*flow.Tweet, error) {

	currentUserId, err := flow.GetUserIDFromContext(ctx)
	if err != nil {
		return &flow.Tweet{}, flow.ErrUnAuthenticated
	}

	input.Sanitize()
	if err := input.Validate(); err != nil {
		return flow.Tweet{}, err
	}

	tweet, err := ts.TweetRepository.Create(ctx, flow.Tweet{
		Body:   input.Body,
		UserID: currentUserId,
	})

	if err != nil {
		return &flow.Tweet{}, err
	}

	return tweet, nil
}

func (s *TweetService) GetID(ctx context.Context, id string) (*flow.Tweet, error) {
	tweet, err := s.TweetRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
}
