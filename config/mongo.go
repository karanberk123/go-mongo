package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var configuration = New()

func NewMongoDatabase() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	mongoPoolMax, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	mongoMaxIdleTime, _ := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))

	option := options.Client().
		ApplyURI(ClientUrl()).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		panic(err)
	}
	database := client.Database("Cluster0")
	fmt.Println(client.Ping(ctx, nil))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func ClientUrl() string {
	clientUrl := "mongodb+srv://karan:1q2w3e4r5t@cluster0.gznog.mongodb.net/?retryWrites=true&w=majority"
	return clientUrl
}
