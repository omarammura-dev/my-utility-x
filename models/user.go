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
	İsVerified bool 
}


func (u *User) Save() error {
	database, ctx, err := db.Init()

	if err != nil {
		return err
	}
	userCollection := database.Database("myutilityx").Collection("users")
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := userCollection.InsertOne(ctx, u)

	u.ID = result.InsertedID.(primitive.ObjectID)
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

func (u User) Update() error {
	database, ctx, err := db.Init()

	if err != nil {
		return err
	}


	if u.ID.IsZero() {
		return errors.New("empty id")
	}
	filter := bson.M{"_id": u.ID}



	_, err = database.Database(os.Getenv("MONGO_DB_NAME")).Collection("users").UpdateOne(ctx,filter,bson.M{"$set": bson.M{"isverified": true}})

	
	return err
}