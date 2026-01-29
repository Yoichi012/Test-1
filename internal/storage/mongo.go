package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/YourUsername/waifu-catcher/internal/config"
)

type MongoStore struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoStore(cfg *config.Config) (*MongoStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.MongoDB)
	return &MongoStore{Client: client, DB: db}, nil
}

func (m *MongoStore) Disconnect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}