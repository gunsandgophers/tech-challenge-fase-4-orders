package database

import (
	"fmt"
	"tech-challenge-fase-1/internal/infra/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDatabase() *mongo.Database {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT)
	clientOption := options.Client().ApplyURI(url)
	client, err := mongo.Connect(clientOption)
	if err != nil {
		panic(err)
	}

	return client.Database(config.DB_NAME)
}
