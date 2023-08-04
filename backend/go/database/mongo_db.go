package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/theLemonday/k8s-todos/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	client *mongo.Client
	db     *mongo.Database
}

var (
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
)

func (mg *MongoInstance) GetCollection(name string) *mongo.Collection {
	return mg.db.Collection(name)
}

func NewConnection(cfg *config.MongodbConfig) (*MongoInstance, error) {
	var mg MongoInstance
	var err error

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/", cfg.Username, cfg.Password, cfg.Host)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	mg.client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	defer cancel()

	mg.db = mg.client.Database(cfg.Database)

	return &mg, nil
}

func (mg MongoInstance) CloseConnection() {
	if err := mg.client.Disconnect(context.TODO()); err != nil {
		log.Panic().Err(err)
	}
}

func (mg MongoInstance) CheckConnection() {
	if err := mg.client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Panic().Err(err)
	}
}
