package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"myutilityx.com/db"
)

type Link struct{
	Name string 
	Url string `binding:"required"`
	CreatedAt time.Time
	User_id int64
}


func (l Link) Save() error {
	database,err,ctx := db.Init()
	if err !=nil {
		return err
	}

	linksCollection := database.Database("myutilityx").Collection("links")
	l.CreatedAt = time.Now()
	linksCollection.InsertOne(ctx,l)
	return err
}

func (l Link) GetAll() (error,[]bson.M){
	database,err,ctx := db.Init()
	if err !=nil {
		return err,nil
	}	
	linksCollection := database.Database("myutilityx").Collection("links")

	cursor,err := linksCollection.Find(ctx,bson.M{})
	if err != nil {
		return err,nil
	}

	var links []bson.M
	if err = cursor.All(ctx,&links); err != nil {
		return err,nil
	}
	return nil,links
}