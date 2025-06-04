package jwt

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/config"
	jwtGo "github.com/lestrrat-go/jwx/jwt"
	"github.com/stretchr/testify/require"
)

var (
	conf         *config.Config
	tokenService *TokenService
)

func TestMain(m *testing.M) {
	config.LoadEnv(".env.test")
	conf = config.New()
	tokenService = NewTokenService(conf)
	os.Exit(m.Run())
}

func TestTokenService_CreateAccessToken(t *testing.T) {
	t.Run("create valid access token", func(t *testing.T) {
		ctx := context.Background()
		user := flow.User{
			ID: "12345",
		}
		token, err := tokenService.CreateAccessToken(ctx, user)
		require.NoError(t, err)

		now = func() time.Time {
			return time.Now()
		}

		tok, err := jwtGo.Parse(
			[]byte(token),
			jwtGo.WithValidate(true),
			jwtGo.WithVerify(signatureType, []byte(conf.JWT.Secret)),
			jwtGo.WithIssuer(conf.JWT.Issuer),
		)

		require.NoError(t, err)
		require.Equal(t, "12345", tok.Subject())
		require.Equal(t, now().Add(flow.AccessTokenLifetime).Unix(), tok.Expiration().Unix())
		teardownTimerNow(t)
	})

}

func TestTokenService_CreateRefreshToken(t *testing.T) {
	t.Run("create valid refresh token", func(t *testing.T) {
		ctx := context.Background()
		user := flow.User{
			ID: "12345",
		}
		token, err := tokenService.CreateRefreshToken(ctx, user, "token-id-12345")
		require.NoError(t, err)

		now = func() time.Time {
			return time.Now()
		}

		tok, err := jwtGo.Parse(
			[]byte(token),
			jwtGo.WithValidate(true),
			jwtGo.WithVerify(signatureType, []byte(conf.JWT.Secret)),
			jwtGo.WithIssuer(conf.JWT.Issuer),
		)

		require.NoError(t, err)
		require.Equal(t, "12345", tok.Subject())
		require.Equal(t, "token-id-12345", tok.JwtID())
		require.Equal(t, now().Add(flow.RefreshTokenLifetime).Unix(), tok.Expiration().Unix())
		teardownTimerNow(t)
	})

}

func TestTokenService_ParseToken(t *testing.T) {
	t.Run("parse valid access token", func(t *testing.T) {
		ctx := context.Background()
		user := flow.User{
			ID: "12345",
		}
		token, err := tokenService.CreateAccessToken(ctx, user)
		require.NoError(t, err)

		tok, err := tokenService.ParseToken(ctx, token)
		require.NoError(t, err)
		require.Equal(t, "12345", tok)
	})

	t.Run("return err when token is invalid", func(t *testing.T) {
		ctx := context.Background()
		_, err := tokenService.ParseToken(ctx, "invalid token")
		require.ErrorIs(t, err, flow.ErrInvalidAccessToken)
	})

	t.Run("return err when token expired", func(t *testing.T) {
		ctx := context.Background()
		user := flow.User{
			ID: "12345",
		}

		now = func() time.Time {
			return time.Now().Add(-flow.AccessTokenLifetime * 5)
		}

		token, err := tokenService.CreateAccessToken(ctx, user)
		require.NoError(t, err)

		_, err = tokenService.ParseToken(ctx, token)
		require.ErrorIs(t, err, flow.ErrInvalidAccessToken)

		teardownTimerNow(t)
	})

}

func TestTokenService_ParseTokenFromRequest(t *testing.T) {
	t.Run("parse valid access token from request", func(t *testing.T) {
		ctx := context.Background()
		user := flow.User{
			ID: "12345",
		}

		req := httptest.NewRequest("GET", "/", nil)

		accessToken, err := tokenService.CreateAccessToken(ctx, user)
		require.NoError(t, err)

		req.Header.Set("Authorization", accessToken)

		token, err := tokenService.ParseTokenFromRequest(ctx, req)
		require.NoError(t, err)

		require.Equal(t, "12345", token)

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		token, err = tokenService.ParseTokenFromRequest(ctx, req)
		require.NoError(t, err)

		require.Equal(t, "12345", token.Sub)
	})

}

func teardownTimerNow(t *testing.T) {
	t.Helper()
	// Reset the now function to its original state
	now = func() time.Time {
		return time.Now()
	}
}
