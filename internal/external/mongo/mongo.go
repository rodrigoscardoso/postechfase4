package mongo

import (
	"context"
	"fmt"
	"log/slog"
	"post-tech-challenge-10soat/internal/infrastructure/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MONGO struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func New(ctx context.Context, config *config.MONGO) (*MONGO, error) {
	// Create connection string with authSource=admin for proper authentication
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	// Create client options with detailed configuration
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetAuth(options.Credential{
		Username:   config.User,
		Password:   config.Password,
		AuthSource: "admin",
		AuthMechanism: "SCRAM-SHA-1",
	})

	// Enable detailed logging
	slog.Info("Attempting MongoDB connection", "uri", maskCredentials(uri))

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		slog.Error("Error connecting to MongoDB", "error", err, "uri", maskCredentials(uri))
		return nil, err
	}

	// Test connection
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(ctxPing, nil); err != nil {
		slog.Error("Error pinging MongoDB", "error", err)
		return nil, err
	}

	slog.Info("Successfully connected to MongoDB", "database", config.Name)

	return &MONGO{
		Client:   client,
		Database: client.Database(config.Name),
	}, nil
}

func maskCredentials(uri string) string {
	return "mongodb://****:****@${DB_HOST}:${DB_PORT}/${DB_NAME}"
}
