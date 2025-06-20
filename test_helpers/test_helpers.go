package test_helpers

import (
	"context"
	"testing"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/faker"
	"github.com/arisromil/flow/postgres"
	"github.com/stretchr/testify/require"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Truncate(ctx)
	require.NoError(t, err, "failed to truncate database")

}

func CreateUser(ctx context.Context, t *testing.T, userRepo flow.UserRepository) flow.User {
	t.Helper()

	userID, err := userRepo.Create(ctx, flow.User{
		Username: faker.Username(),
		Password: faker.Password,
		Email:    faker.Email(),
	})
	require.NoError(t, err, "failed to create user")

	return userID
}

func createTweet(ctx context.Context, t *testing.T, tweetRepo flow.TweetRepository, forUser string) flow.Tweet {
	t.Helper()

	tweet, err := tweetRepo.CreateTweet(ctx, flow.Tweet{
		Body:   faker.RandomString(100),
		UserID: forUser,
	})
	require.NoError(t, err, "failed to create tweet")

	return tweet
}

func LoginUser(ctx context.Context, t *testing.T, user flow.User) context.Context {
	t.Helper()

	return flow.PutUserIdToContext(ctx, user.ID)
}
