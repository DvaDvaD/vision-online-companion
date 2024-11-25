package main

import (
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(args []string, db *sql.DB) {
	if len(args) < 1 {
		log.Info().Msgf("%+v %+v", args, db)
		log.Fatal().Msg("Please insert an action (up, down, reset)")
		return
	}

	log.Info().Msgf("Starting database migration: %s", args[0])
	goose.SetDialect("postgres")
	goose.SetBaseFS(migrationFiles)

	action := args[0]

	switch action {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatal().Err(err).Msg("Failed to run up migrations")
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatal().Err(err).Msg("Failed to run down migrations")
		}
	case "reset":
		if err := goose.Reset(db, "migrations"); err != nil {
			log.Fatal().Err(err).Msg("Failed to reset migrations")
		}
	default:
		log.Fatal().Msgf("Invalid action. Please use 'up', 'down', or 'reset'. Currently use: %s", action)
	}

	log.Info().Msg("Migrations ran successfully")
}
