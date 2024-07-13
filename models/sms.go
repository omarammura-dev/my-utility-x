package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
