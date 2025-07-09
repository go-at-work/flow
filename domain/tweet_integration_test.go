//go:build integration

package domain

import (
	"context"
	"testing"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/faker"
	"github.com/arisromil/flow/test_helpers"
	"github.com/stretchr/testify/require"
)

func TestIntegrationTweetService_Create(t *testing.T) {
	t.Run("user not authorized to Create Tweet", func(t *testing.T) {
		ctx := context.Background()

		_, err := tweetService.CreateTweet(ctx, flow.CreateTweetRequest{
			Body: "This is a test tweet",
		})

		require.Error(t, err, "Expected error when creating tweet without authorization")
	})

	t.Run("user authorized to Create Tweet", func(t *testing.T) {

		ctx := context.Background()

		defer test_helpers.TearDownDB(ctx, t, db)

		currentUserId := test_helpers.CreateUser(ctx, t, userRepo)

		ctx = test_helpers.LoginUser(ctx, t, currentUserId)

		input := flow.CreateTweetRequest{
			Body: faker.RandomString(100),
		}

		tweet, err := tweetService.CreateTweet(ctx, input)
		require.NoError(t, err, "Expected no error when creating tweet with valid input")

		require.NotEmpty(t, tweet.ID, "Expected tweet ID to be generated")
		require.Equal(t, input.Body, tweet.Body, "Expected tweet body to match input")
		require.Equal(t, currentUserId, tweet.UserID, "Expected tweet to be associated with the correct user")
		requore.NotEmpty(t, tweet.CreatedAt, "Expected tweet to have a creation timestamp")

	})
}

func TestIntegrationTweetService_GetByID(t *testing.T) {
	t.Run("can get a tweet by id", func(t *testing.T) {
		ctx := context.Background

		defer test_helpers.TearDownDB(ctx, t, db)

		currentUserId := test_helpers.CreateUser(ctx, t, userRepo)
		existingTweet := test_helpers.CreateTweet(ctx, t, tweetRepo, currentUserId)

		tweet, err := tweetService.GetID(ctx, existingTweet.ID)
		require.NoError(t, err, "Expected no error when getting tweet by ID")
		require.Equal(t, existingTweet.Id, tweet.ID, "Expected tweet ID to match the one created")
		require.Equal(t, existingTweet.Body, tweet.Body, "Expected tweet body to match the one created")

	})

	t.Run("cannot get a tweet by non-existing id", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TearDownDB(ctx, t, db)

		_, err := tweetService.GetID(ctx, faker.UUID())
		require.ErrorIs(t, err, "Expected error when getting tweet by non-existing ID")
		
	}	

    t.Run("return error invalud uuid", func(t *testing.T) {
	ctx := context.Background()

		defer test_helpers.TearDownDB(ctx, t, db)

		_, err := tweetService.GetID(ctx, "123")
		require.ErrorIs(t, err, flow.ErrInvalidUUID)
    }

}


func TestIntegrationTweetService_All(t *testing.T) {
	t.Run("return all tweets", func(t *testing.T) {
		ctx := context.Background

		defer test_helpers.TearDownDB(ctx, t, db)

		currentUserId := test_helpers.CreateUser(ctx, t, userRepo)

		test_helpers.CreateTweet(ctx, t, tweetRepo, currentUserId)
		test_helpers.CreateTweet(ctx, t, tweetRepo, currentUserId)
		test_helpers.CreateTweet(ctx, t, tweetRepo, currentUserId)

		tweets, err := tweetService.All(ctx)
		require.NoError(t, err, "Expected no error when getting tweet by ID")

		require.Len(t, tweets, 3, "Expected to retrieve 3 tweets")
		require.Equal(t, existingTweet.Body, tweet.Body, "Expected tweet body to match the one created")

	})


}
