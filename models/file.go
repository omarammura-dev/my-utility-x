package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FileName  string             `bson:"fileName" json:"fileName"`
  ContentType string                        `bson:"contentType" json:"contentType"`
	FileId    string             `bson:"fileId" binding:"required" json:"fileId"`
	Locked    bool               `bson:"locked" json:"locked"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UserId    primitive.ObjectID `bson:"userId" json:"userId"`
}
