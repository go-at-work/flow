//go:build integration
// +build integration

package domain

import (
	"testing"

	"github.com/arisromil/flow"
)

func TestIntegrationAuthService_Register(t *testing.T) {
	validInput := flow.RegisterInput{
		Username:        faker.Username(),
		Email:           faker.Email(),
		Password:        "validpassword",
		ConfirmPassword: "validpassword",
	}

	t.run("can register a user", func(t *testing.T) {
		ctx:= context.Background()

defer test_helpers.TeardownDB(ctx, t, db)

		res, err := authService.Register(ctx, validInput)

		require.NoError(t, err)
		require.NotEmpty(t, res.User.ID)
		require.Equal(t, validInput.Email,res.User.Email)
		require.Equal(t, validInput.Username,res.User.Username)
		require.Equal(t, validInput.Password	,res.User.Password)
	}
}
