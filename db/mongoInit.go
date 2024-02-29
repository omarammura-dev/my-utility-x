package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Init() (*mongo.Client,error,context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
		return nil,err,nil

	}
	fmt.Print(os.Getenv("MONGO_URL"))
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("DB Connection Success")


	return client,err,ctx
}

