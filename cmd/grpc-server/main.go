package main

import (
	"flag"
	"github.com/Alex1472/ozon-category-service/internal/service/category"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/Alex1472/ozon-category-service/internal/config"
	"github.com/Alex1472/ozon-category-service/internal/server"
	"github.com/Alex1472/ozon-category-service/internal/service/category/repository"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	categoryRepository := category_repository.New()
	categoryService := category.New(categoryRepository)

	if err := server.NewGrpcServer(categoryService).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
