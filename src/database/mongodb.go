package database

import (
	"clinic-api/src/configs"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
	config, _ := configs.LoadServerConfig(".")
	dsn := config.MongoConnectionString

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	return client
}
