package main

import (
	"net/http"

	"github.com/arisromil/flow"
)

func authMiddleware(authTokenService flow.AuthTokenService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := authTokenService.ParseTokenFromRequest(ctx, r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// add the user id to the context
			ctx = flow.P.WithValue(ctx, "userID", token.Sub)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
