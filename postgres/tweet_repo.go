package postgres

import (
	"context"
	"fmt"

	"github.com/arisromil/flow"
	"github.com/georgysavva/scany/pgxscan"
)

type TweetRepository struct {
	DB *DB
}

func NewTweetRepo(db *DB) *TweetRepository {
	return &TweetRepository{
		DB: db,
	}
}

func (s *TweetRepository) All(ctx context.Context) ([]*flow.Tweet, error) {
	panic("not implemented")
}

func (tr *TweetRepository) Create(ctx context.Context, tweet flow.Tweet) (*flow.Tweet, error) {
	tx, err := tr.DB.Pool.Begin(ctx)
	if err != nil {
		return flow.Tweet{}, fmt.Errorf("error starting transaction: %v", err)
	}

	defer tx.Rollback(ctx)

	tweet, err = createTweet(ctx, tx, tweet)
	if err != nil {
		return &flow.Tweet{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return flow.Tweet{}, fmt.Errorf("error commit: %v", err)
	}

	return &tweet, nil

}

func createTweet(ctx context.Context, tx pgxscan.Querier, tweet flow.Tweet) (flow.Tweet, error) {
	query := `INSERT INTO tweets (body, user_id, password) VALUES ($1, $2) RETURNING *;`
	t := flow.Tweet{}
	if err := pgxscan.Get(ctx, tx, &t, query, tweet.Body, tweet.UserID); err != nil {
		return flow.Tweet{}, fmt.Errorf("Error creating user: %w", err)
	}

	return t, nil
}

func (s *TweetRepository) GetID(ctx context.Context, id string) (*flow.Tweet, error) {
	panic("not implemented")
}
