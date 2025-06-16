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

func LoginUser(ctx context.Context, t *testing.T, userID flow.User) context.Context {
	t.Helper()

	token, err := userRepo.Login(ctx, userID)
	require.NoError(t, err, "failed to login user")

	return flow.PutUserIdToContext(ctx, token.Sub)
}
