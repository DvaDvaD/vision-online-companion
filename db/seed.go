package main

import (
	"database/sql"
	"embed"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

//go:embed seeds/*.sql
var seedFiles embed.FS

func SeedFromFiles(args []string, db *sql.DB) {
	log.Info().Msg("Starting database seed")

	files, err := seedFiles.ReadDir("seeds")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read seeds directory")
		return
	}

	if len(files) == 0 {
		log.Info().Msg("No files to seed")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		filePath := filepath.Join("seeds", fileName)

		content, err := seedFiles.ReadFile(filePath)
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to read file %s", filePath)
			return
		}

		queries := strings.Split(string(content), ";")
		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" {
				continue
			}
			if _, err := db.Exec(query); err != nil {
				log.Fatal().Err(err).Msgf("Could not execute SQL query from file %s with query: %v", fileName, query)
				return
			}
		}
		log.Info().Msgf("%s seeded successfully", fileName)
	}

	log.Info().Msg("Seeder ran successfully")
}