package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arisromil/flow/config"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := porstgres.NewDB(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("migrated database")

}
