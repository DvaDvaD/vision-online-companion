package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/portierglobal/vision-online-companion/database/data"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type application struct {
	logger         *zerolog.Logger
	queries        *data.Queries
	shutdownCancel context.CancelFunc
}

func init() {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("No .env file found, relying on environment variables")
	}
}

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.AutomaticEnv()          // read environment variables that match

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("Connecting to database...")

	db, err := pgx.Connect(context.Background(), viper.GetString("DB_DSN"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer db.Close(context.Background())

	app := &application{
		logger:  &log.Logger,
		queries: data.New(db),
	}

	err = app.serve()
	if err != nil {
		app.logger.Fatal().Err(err).Msg("shutting down")
	}
}
