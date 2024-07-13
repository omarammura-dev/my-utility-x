package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myutilityx.com/db"
)

type SMS struct {
	ID   primitive.ObjectID `bson:"_id"`
	Code string             `json:"code"`
}

func (s *SMS) Save() error {
	database, ctx, err := db.Init()

	if err != nil {
		return err
	}
	smsCollection := database.Database("myutilityx").Collection("sms")
	s.ID = primitive.NewObjectID()
	result, err := smsCollection.InsertOne(ctx, s)
	if err != nil {
		return err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		s.ID = oid
	} else {
		return errors.New("type assertion of InsertedID to primitive.ObjectID failed")
	}
	return err
}
func (s *SMS) Get() error {
	database, ctx, err := db.Init()
	if err != nil {
		return err
	}
	coll := database.Database("myutilityx").Collection("sms")
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	if err = coll.FindOne(ctx, bson.M{}, opts).Decode(&s); err != nil {
		return err
	}
	return nil
}

func (s *SMS) Delete() error {
	database, ctx, err := db.Init()
	if err != nil {
		return err
	}

	coll := database.Database("myutilityx").Collection("sms")
	coll.DeleteMany(ctx, bson.M{"code": s.Code})
	return nil
}
