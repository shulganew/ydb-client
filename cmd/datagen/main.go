package main

import (
	"context"
	"fmt"
	"path"
	"time"
	"ydb-client/internal/config"
	"ydb-client/internal/storage"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/sugar"
)

func main() {
	cfg := config.GetConfig()

	if cfg.Series <= 0 || cfg.Series > 100_000_000 {
		panic(fmt.Errorf("expected employees count to be 0 <= count <= 1_000_000, got: %d", cfg.Series))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("connection sring: ", cfg.DSN)
	db, err := ydb.Open(ctx, cfg.DSN)
	if err != nil {
		panic(fmt.Errorf("connect error: %w", err))
	}
	defer func() { _ = db.Close(ctx) }()

	prefix := path.Join(db.Name(), "native/query")

	fmt.Println("Prifix: ", prefix)

	// Clean database.
	err = sugar.RemoveRecursive(ctx, db, prefix)
	if err != nil {
		panic(err)
	}

	err = storage.CreateTables(ctx, db.Table(), prefix)
	if err != nil {
		panic(fmt.Errorf("create tables error: %w", err))
	}

	// Tablse service options.
	err = storage.DescribeTableOptions(ctx, db.Table())
	if err != nil {
		panic(fmt.Errorf("describe table options error: %w", err))
	}
	// Describe created tables.
	err = storage.DescribeTable(ctx, db.Table(), path.Join(prefix, "series"))
	if err != nil {
		panic(fmt.Errorf("create tables error: %w", err))
	}

	err = storage.DescribeTable(ctx, db.Table(), path.Join(prefix, "seasons"))
	if err != nil {
		panic(fmt.Errorf("create tables error: %w", err))
	}

	err = storage.DescribeTable(ctx, db.Table(), path.Join(prefix, "episodes"))
	if err != nil {
		panic(fmt.Errorf("create tables error: %w", err))
	}

	// Generate table data.
	err = storage.FacialillTablesWithData(ctx, cfg, db.Table(), prefix)
	if err != nil {
		panic(fmt.Errorf("fill tables with data error: %w", err))
	}

}
