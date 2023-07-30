package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hugeman/todolist/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var Ctx context.Context
var Client *mongo.Client
var Cancel context.CancelFunc

func InitalDatabase() error {
	database, ctx, client, cancel, err := newPool(
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)
	if err != nil {
		return err
	}

	Database = database
	Ctx = ctx
	Client = client
	Cancel = cancel

	return nil
}

func newPool(host string, port string, name string) (*mongo.Database, context.Context, *mongo.Client, context.CancelFunc, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s",
		host,
		port,
	))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	db := client.Database(name)

	return db, ctx, client, cancel, nil
}
