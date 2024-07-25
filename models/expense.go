package models

import (
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myutilityx.com/db"
)

type Expense struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ExpenseName string             `bson:"expenseName" validate:"required"`
	ExpenseDate time.Time          `bson:"expenseDate" validate:"required"`
	UserId      primitive.ObjectID `bson:"userId" validate:"required"`
	ExpenseType string             `bson:"expenseType" validate:"required"`
	Price       float64            `bson:"price" validate:"required"`
}

func (e *Expense) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

func (e *Expense) Save() error {

	if err := e.Validate(); err != nil {
		return err
	}

	database, ctx, err := db.Init()

	if err != nil {
		return err
	}

	expenseCollection := database.Database(os.Getenv("MONGO_DB_NAME")).Collection("expense")

	_, err = expenseCollection.InsertOne(ctx, e)

	return err
}

func UpdateExpense(payload map[string]interface{}) error {
	database, ctx, err := db.Init()

	if err != nil {
		return err
	}
	update := bson.M{
		"$set": payload,
	}

	_, err = database.Database(os.Getenv("MONGO_DB_NAME")).Collection("expense").UpdateOne(ctx, bson.M{"_id": payload["_id"]}, update)
	return err
}

func (e *Expense) GetAllExpenses(userId primitive.ObjectID) ([]bson.M, error) {
	database, ctx, err := db.Init()

	if err != nil {
		return nil, err
	}

	expenseCollection := database.Database(os.Getenv("MONGO_DB_NAME")).Collection("expense")

	cursor, err := expenseCollection.Find(ctx, bson.M{"userid": userId})
	if err != nil {
		return nil, err
	}

	var expenses []bson.M

	if err = cursor.All(ctx, &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}
