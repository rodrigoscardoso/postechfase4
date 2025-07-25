package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "post-tech-challenge-10soat/docs"
	router "post-tech-challenge-10soat/internal/delivery/http"
	"post-tech-challenge-10soat/internal/external/mongo"
	"post-tech-challenge-10soat/internal/external/postgres"
	"post-tech-challenge-10soat/internal/infrastructure/config"
	dependency "post-tech-challenge-10soat/internal/infrastructure/di"
	"post-tech-challenge-10soat/internal/infrastructure/logger"
)

//	@title			POS-Tech API
//	@version		1.0
//	@description	API em Go para o desafio na pos-tech fiap de Software Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	conf, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	logger.Set(conf.App)
	slog.Info("Starting the application", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()
	db, errPostgres := postgres.New(ctx, conf.DB)

	if errPostgres != nil {
		slog.Error("Error initializing database connection", "error", errPostgres)
		os.Exit(1)
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, errMongo := mongo.New(ctxMongo, conf.MONGO)

	if errMongo != nil {
		slog.Error("Error initializing database connection", "error", errMongo)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "MONGO", conf.MONGO.Connection)

	slog.Info("Successfully connected to the database", "DB", conf.DB.Connection)

	errPostgresMigrate := db.Migrate()
	if errPostgresMigrate != nil {
		slog.Error("Error migrating database", "error", errPostgresMigrate)
		os.Exit(1)
	}

	defer db.Close()

	if (errMongo != nil) || (errPostgresMigrate != nil || errPostgres != nil) {
		slog.Error("Error initializing database connection", "error", errMongo)
		os.Exit(1)
	}

	// di
	healthHandler, clientHandler, productHandler, orderHandler := dependency.Setup(conf.App, db, mongo)

	router, err := router.NewRouter(
		conf.HTTP,
		healthHandler,
		clientHandler,
		productHandler,
		orderHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddress := fmt.Sprintf("%s:%s", conf.HTTP.URL, conf.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddress)
	err = router.Run(listenAddress)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
