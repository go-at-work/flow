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
