package bd

import (
	"context"
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func ConectToBd() *bun.DB {
	configcon := pgdriver.NewConnector(
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithInsecure(true),
		pgdriver.WithDatabase("electricity"),
		pgdriver.WithPassword("docker"),
		pgdriver.WithUser("test"),
	)
	sqlcon := sql.OpenDB(configcon)
	db := bun.NewDB(sqlcon, pgdialect.New())
	return db
}

func CloseToBd(db *bun.DB) {
	db.Close()
}

func Migrations(ctx context.Context, db bun.DB) {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}
}
