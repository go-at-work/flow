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
	authService flow.AuthService
	userRepo    flow.UserRepository
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

	authService.NewAuthService(userRepo)
	os.Exit(t.Run())

	// This is a placeholder for the main function.
	// In a real-world scenario, this would be the entry point of the application.
}
