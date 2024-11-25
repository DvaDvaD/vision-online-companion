package main

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to read the config file")
	}

	dsn := viper.GetString("database.dsn")
	log.Info().Msgf("Connecting to the database %+v", dsn)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping the database")
	}
	defer db.Close()

	rootCmd := &cobra.Command{Use: "app"}

	migrateCmd := &cobra.Command{
		Use:   "migrate [action]",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			Migrate(args, db)
		},
	}

	seedCmd := &cobra.Command{
		Use:   "seed",
		Short: "Seed the database",
		Run: func(cmd *cobra.Command, args []string) {
			SeedFromFiles(args, db)
		},
	}

	rootCmd.AddCommand(migrateCmd, seedCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
