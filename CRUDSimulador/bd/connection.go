// Package bd contains all instructions that call the database
package bd

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// ConectToBd Create the connection with the database
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

// CloseToBd Close the connection with the database
func CloseToBd(db *bun.DB) {
	if err := db.Close(); err != nil {
		log.Print(err)
	}
}

// Migrations method that create the migration in the database
func Migrations(db bun.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}
}
