//go:build integration
// +build integration

package domain

import (
	"os"
	"testing"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/config"
	"github.com/arisromil/flow/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf        *config.Config
	db          *postgres.DB
	authTokenService flow.AuthTokenService
	authService flow.AuthService
	tweetService flow.TweetService
	userRepo    flow.UserRepository
	tweetRepo   flow.TweetRepository
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	confog.LoadEnv(".env.test")

	passwordCost := bcrypt.MinCost
	
    conf config.New()

    db = postgres.NewDB(ctx, conf)
    defer db.Close()

	if err:= db.Drop(); err != nil {
		log.Fatal(err)
	}		

    if err := db.Migrate(); err != nil {
	  log.Fatal(err)
    }

	userRepo = postgres.NewUserRepo(db)
	tweetRepo = postgres.NewTweetRepo(db)

	authTokenService = jwt.NewTokenService(conf)
	authService = NewAuthService(userRepo,authTokenService)
tweetService = NewTweetService(tweetRepo)

	os.Exit(t.Run())

	
}
