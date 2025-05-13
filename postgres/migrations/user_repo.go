package postgres

import (
	"context"
	"fmt"

	"github.com/arisromil/flow"
	"github.com/georgysavva/scany/pgxscan"
)

type UserRepo struct {
	DB *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (ur *UserRepo) Create(ctx context.Context, user flow.User) (flow.User, error) {

	tx, err := ur.DB.Pool.Begin(ctx)
	if err != nil {
		return flow.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	user, err = createUser(ctx, tx, user)
	if err != nil {
		return flow.User{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return flow.User{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user, nil

}

func createUser(ctx context.Context, tx pgxscan.Querier, user flow.User) (flow.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *;`
	u := flow.User{}
	if err := pgxscan.Get(ctx, tx, &u, query, user.Username, user.Email, user.Password); err != nil {
		return flow.User{}, fmt.Errorf("Error creating user: %w", err)
	}

	return u, nil
}

func (ur *UserRepo) GetbyUsername(ctx context.Context, username string) (flow.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`
	u := flow.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return flow.User{}, flow.ErrNotFound
		}

		return flow.User{}, fmt.Errorf("Error getting user by username: %w", err)

	}
	return u, nil
}

func (ur *UserRepo) GetbyEmail(ctx context.Context, email string) (flow.User, error) {
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	u := flow.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return flow.User{}, flow.ErrNotFound
		}

		return flow.User{}, fmt.Errorf("Error getting user by username: %w", err)

	}
	return u, nil
}
