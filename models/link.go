package models

import (
	"context"
	"time"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myutilityx.com/db"
)

type Link struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Url       string             `bson:"url" binding:"required" json:"url"`
	ShortUrl  string             `bson:"shortUrl" json:"shortUrl"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	Clicks    int64              `bson:"clicks" json:"clicks"`
	UserId    primitive.ObjectID `bson:"userId" json:"userId"`
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

	database, ctx, err := db.Init()
	if err != nil {
		return err
	}

	linksCollection := database.Database("myutilityx").Collection("links")
	l.CreatedAt = time.Now()
	_, err = linksCollection.InsertOne(ctx, l)
	return err
}
func (l Link) GetAll(userID primitive.ObjectID) ([]bson.M, error) {

	database, ctx, err := db.Init()
	if err != nil {
		return nil, err
	}

	linksCollection := database.Database("myutilityx").Collection("links")

	cursor, err := linksCollection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}

	var links []bson.M
	if err = cursor.All(ctx, &links); err != nil {
		return nil, err
	}
	return links, nil
}

func FindById(id primitive.ObjectID) (*Link, error) {
	database, ctx, err := db.Init()
	if err != nil {
		return nil, err
	}
	linksCollection := database.Database("myutilityx").Collection("links")

	var link Link
	err = linksCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func GetSingleAndIncreaseClicks(shortUrl string) (*Link, error) {
	database, _, err := db.Init()
	if err != nil {
		return nil, err
	}
	linksCollection := database.Database("myutilityx").Collection("links")

	var results Link
	err = linksCollection.FindOne(context.TODO(), bson.M{"shortUrl": shortUrl}).Decode(&results)
	if err != nil {
		return nil, err
	}
	_, err = linksCollection.UpdateOne(context.Background(), bson.M{"shortUrl": shortUrl}, bson.M{"$set": bson.M{"clicks": results.Clicks + 1}})
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (l Link) Delete() error {
	database, ctx, err := db.Init()
	if err != nil {
		return err
	}
	linksCollection := database.Database("myutilityx").Collection("links")

	_, err = linksCollection.DeleteOne(ctx, bson.M{"_id": l.Id})

	return err
}
