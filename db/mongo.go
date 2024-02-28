package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myutilityx.com/models"
)

var collection *mongo.Collection
var link models.Link

func Init(collectionName string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
		return err
	}
	fmt.Print(os.Getenv("MONGO_URL"))
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("DB Connection Success")

	collection = client.Database("myutilityx").Collection(collectionName)
	fmt.Print("Collection Created", collectionName)
	return err
}




func Addl(ctx *gin.Context) {

	link = models.Link{
		Name:      "Example Link",
		Url:       "https://example.com",
		CreatedAt: time.Now(),
		User_id:   1,
	}


	insertion, err := collection.InsertOne(context.Background(), link)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("inserted !!! ", insertion)

	ctx.JSON(http.StatusOK,gin.H{"message":"inserted!"})
}
