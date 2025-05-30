package jwt

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/config"
	"github.com/lestrrat-go/jwx/jwa"
	jwtGo "github.com/lestrrat-go/jwx/jwt"
)

var signatureType = jwa.HS256

type TokenService struct {
	Conf *config.Config
}

func NewTokenService(conf *config.Config) *TokenService {
	return &TokenService{
		Conf: conf,
	}
}

func (ts *TokenService) ParseTokenFromRequest(ctx context.Context, r *http.Request) (flow.AuthToken, error) {
	token, err := jwtGo.ParseRequest(
		r,
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)))

	if err != nil {
		return flow.AuthToken{}, flow.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func buildToken(token jwtGo.Token) flow.AuthToken {
	return flow.AuthToken{
		ID:  token.JwtID(),
		Sub: token.Subject(),
	}

}

func (ts *TokenService) ParseToken(ctx context.Context, payload string) (flow.AuthToken, error) {
	token, err := jwtGo.Parse(
		[]byte(payload),
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)))

	if err != nil {
		return flow.AuthToken{}, flow.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func (ts *TokenService) CreateRefreshToken(ctx context.Context, user flow.User, tokenId string) (string, error) {
	t := jwtGo.New()

	if err := setDefaultToken(t, user, flow.RefreshTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	if err := t.Set(jwtGo.JwtIDKey, tokenId); err != nil {
		return "", fmt.Errorf("failed to set jwt id")
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return string(token), nil
}

func (ts *TokenService) CreateAccessToken(ctx context.Context, user flow.User) (string, error) {
	t := jwtGo.New()

	if err := setDefaultToken(t, user, flow.AccessTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return string(token), nil
}

func setDefaultToken(t jwtGo.Token, user flow.User, lifetime time.Duration, conf *config.Config) error {
	if err := t.Set(jwtGo.SubjectKey, user.ID); err != nil {
		return fmt.Errorf("failed to set subject key: %w", err)
	}

	if err := t.Set(jwtGo.IssuerKey, conf.JWT.Issuer); err != nil {
		return fmt.Errorf("failed to set issuer key: %w", err)
	}

	if err := t.Set(jwtGo.IssuedAtKey, time.Now().Unix()); err != nil {
		return fmt.Errorf("failed to set issued at key: %w", err)
	}

	if err := t.Set(jwtGo.ExpirationKey, time.Now().Add(lifetime).Unix()); err != nil {
		return fmt.Errorf("failed to set expired at key: %w", err)
	}

	return nil
}
