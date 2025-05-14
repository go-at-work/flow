package test_helpers

import (
	"context"
	"testing"

	"github.com/arisromil/flow/postgres"
	"github.com/stretchr/testify/require"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Truncate(ctx)
	require.NoError(t, err, "failed to truncate database")

}
