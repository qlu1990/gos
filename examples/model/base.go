package model

import (
	"context"
	"time"

	"github.com/qlu1990/gos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient control mongodb
type MongoClient struct {
	Ctx     context.Context
	Client  *mongo.Client
	Timeout time.Duration
}

// Mongo variable
var Mongo = MongoClient{
	Ctx:     context.Background(),
	Timeout: 10 * time.Second,
}

// SetUp set up Mongo variable
func SetUp(uri string) {
	ctx, _ := context.WithTimeout(Mongo.Ctx, Mongo.Timeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		gos.Fatal(err)
	}
	Mongo.Client = client
}
