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
	tweets, err := getAllTweets(ctx, s.DB.Pool)
	if err != nil {
		return nil, err
	}
	result := make([]*flow.Tweet, len(tweets))
	for i := range tweets {
		result[i] = &tweets[i]
	}
	return result, nil
}

func getAllTweets(ctx context.Context, q pgxscan.Querier) ([]flow.Tweet, error) {
	query := `SELECT * FROM tweets;`

	var tweets []flow.Tweet
	if err := pgxscan.Select(ctx, q, &tweets, query); err != nil {
		if pgxscan.NotFound(err) {
			return nil, flow.ErrNotFound
		}
		return nil, fmt.Errorf("error getting all tweets: %w", err)
	}

	return tweets, nil
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

func (s *TweetRepository) GetById(ctx context.Context, id string) (*flow.Tweet, error) {
	tweet, err := getTweetByID(ctx, s.DB.Pool, id)
	if err != nil {
		return nil, err
	}
	return &tweet, nil
}

func getTweetByID(ctx context.Context, q pgxscan.Querier, id string) (flow.Tweet, error) {
	query := `SELECT * from tweets where id = $1 LIMIT 1;`

	t := flow.Tweet{}

	if err := pgxscan.Get(ctx, q, &t, query, id); err != nil {
		if pgxscan.NotFound(err) {
			return flow.Tweet{}, flow.ErrNotFound
		}
		return flow.Tweet{}, fmt.Errorf("error getting tweet by id: %w", err)
	}

	return t, nil
}
