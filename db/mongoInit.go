package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() (*mongo.Client, context.Context,error) {

	fmt.Print(os.Getenv("MONGO_URL"))
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	client, err := mongo.Connect(ctx, opts)
	
	if err != nil {
		log.Fatal(err)
	}
	return client,ctx,err
}

