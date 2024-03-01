package models

import (
	"context"
	"time"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"myutilityx.com/db"
)

type Link struct {
	Name      string
	Url       string `binding:"required"`
	ShortUrl  string
	CreatedAt time.Time
	Clicks    int64
	User_id   int64
}

func InitLink() (*Link, error) {
	link := &Link{}
	id, err := shortid.Generate()
	if err != nil {
		return &Link{}, err
	}
	link.CreatedAt = time.Now()
	link.Clicks = 0
	link.ShortUrl = "U" + id

	return link, nil
}

func (l *Link) Save() error {

	database, err, ctx := db.Init()
	if err != nil {
		return err
	}

	linksCollection := database.Database("myutilityx").Collection("links")
	l.CreatedAt = time.Now()
	linksCollection.InsertOne(ctx, l)
	return err
}
func (l Link) GetAll() ([]bson.M, error) {
	database, err, ctx := db.Init()
	if err != nil {
		return nil, err
	}
	linksCollection := database.Database("myutilityx").Collection("links")

	cursor, err := linksCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var links []bson.M
	if err = cursor.All(ctx, &links); err != nil {
		return nil, err
	}
	return links, nil
}

func GetSingleAndIncreaseClicks(shortUrl string) (*Link, error) {
	database, err, _ := db.Init()
	if err != nil {
		return nil, err
	}
	linksCollection := database.Database("myutilityx").Collection("links")

	var results Link
	err = linksCollection.FindOne(context.TODO(), bson.M{"shorturl": shortUrl}).Decode(&results)
	if err != nil {
		return nil, err
	}
	_, err = linksCollection.UpdateOne(context.Background(), bson.M{"shorturl": shortUrl}, bson.M{"$set":bson.M{"clicks":results.Clicks+1}})
	if err != nil {
		return nil, err
	}
	

	return &results, nil
}
