package postgres

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/arisromil/flow/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
	conf *config.Config
}

func NewDB(ctx context.Context, conf *config.Config) *DB {
	dbConfig, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatalf("failed to parse database config: %v", err)
	}

	pool, errctx := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", errctx)
	}
	db := &DB{
		Pool: pool,
		conf: conf}

	db.Ping(ctx)
	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping postgres database: %v", err)
	}
	log.Println("connected to postgres database")
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

func (db *DB) Migrate() error {
	_, b, _, _ := runtime.Caller(0)
	migrationPath := fmt.Sprintf("file:///%s/migrations", filepath.Dir(b))
	m, err := migrate.New("./postgres/migrations", db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migration: %w", err)
	}

	log.Println("migrations applied successfully")
	return nil
}

func (db *DB) Drop() error {
	_, b, _, _ := runtime.Caller(0)
	migrationPath := fmt.Sprintf("file:///%s/migrations", filepath.Dir(b))
	m, err := migrate.New("./postgres/migrations", db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to drop: %w", err)
	}

	log.Println("migrations applied successfully")
	return nil
}

func (db *DB) Truncate(ctx context.Context) error {
	if _, err := db.Pool.Exec(ctx, "DELETE from users ;"); err != nil {
		return fmt.Errorf("failed to truncate users table: %w", err)
	}
	log.Println("users table truncated successfully")
	return nil
}
