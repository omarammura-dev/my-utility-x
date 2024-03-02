package models

import (
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
	"myutilityx.com/db"
	"myutilityx.com/utils"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	database, ctx, err := db.Init()

	if err != nil {
		return err
	}
	userCollection := database.Database("myutilityx").Collection("users")
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = userCollection.InsertOne(ctx, u)

	return err
}

func (u *User) ValidateCredintials() error {

	database, ctx, err := db.Init()

	if err != nil {
		return err
	}

	filter := bson.M{"email": u.Email}
	projection := bson.M{"password": 1}

	var result struct {
		ID       primitive.ObjectID `bson:"_id"`
		Password string
	}

	err = database.Database(os.Getenv("MONGO_DB_NAME")).Collection("users").FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)

	if err != nil {
		return err
	}

	isValid := utils.CheckPasswordHash(u.Password, result.Password)
	u.ID = result.ID
	if !isValid {
		return errors.New("invalid creditionals")
	}
	return nil
}
