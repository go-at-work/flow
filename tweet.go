package flow

import (
	"context"
	"fmt"
	"strings"
	"time"
)

var (
	TweetMinlength = 2
	TweetMaxlength = 30
	ErrValidation  = fmt.Errorf("validation error")
)

type CreateTweetRequest struct {
	Body string
}

func (in *CreateTweetRequest) Sanitize() {
	in.Body = strings.TrimSpace(in.Body)
}

func (in CreateTweetRequest) Validate() error {
	if len(in.Body) < TweetMinlength {
		return fmt.Errorf("%w: not long enough, at least (%d)", ErrValidation, TweetMinlength)
	}

	if len(in.Body) < TweetMaxlength {
		return fmt.Errorf("%w: body too long, max of  (%d)", ErrValidation, TweetMaxlength)
	}

	return nil

}

type Tweet struct {
	ID        string
	Body      string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TweetService interface {
	All(ctx context.Context) ([]Tweet, error)
	CreateTweet(ctx context.Context, req CreateTweetRequest) (Tweet, error)
	GetById(ctx context.Context, id string) (Tweet, error)
}

type TweetRepository interface {
	All(ctx context.Context) ([]Tweet, error)
	CreateTweet(ctx context.Context, tweet Tweet) (Tweet, error)
	GetById(ctx context.Context, id string) (Tweet, error)
}
