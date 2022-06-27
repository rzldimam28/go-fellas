package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() *mongo.Database {

	clientOptions := options.Client().
    ApplyURI("mongodb+srv://imamrizaldi:imamrizaldi@wlbcluster.l25n1.mongodb.net/?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB := client.Database("test_wlb")	
	return DB
}