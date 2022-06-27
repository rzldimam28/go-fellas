package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rzldimam28/wlb-test/model/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() *mongo.Database {

	err := godotenv.Load()
	helper.PanicIfError(err)

	clientOptions := options.Client().
    ApplyURI(os.Getenv("MONGO_ATLAS_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB := client.Database(os.Getenv("DB_NAME"))	
	return DB
}