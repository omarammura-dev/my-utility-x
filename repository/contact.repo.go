package repository

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myutilityx.com/db"
	"myutilityx.com/graph/model"
)

type ContactRepository interface {
	CreateContact(contact *model.ContactInput) (*model.Contact, error)
	GetContact(id string) (*model.Contact, error)
}

func NewContactRepository() ContactRepository {
	client, _, _ := db.Init()

	return &database{
		client: client,
	}
}

func (d *database) CreateContact(contactInput *model.ContactInput) (*model.Contact, error) {
	client, _, _ := db.Init()

	var contact model.Contact
	if err := mapstructure.Decode(contactInput, &contact); err != nil {
		return nil, err
	}
	contact.ID = primitive.NewObjectID().Hex()
	contact.CreatedAt = timePtr(time.Now())

	contactCollection := client.Database("myutilityx").Collection("contacts")

	_, err := contactCollection.InsertOne(context.TODO(), &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (d *database) GetContact(id string) (*model.Contact, error) {
	client, _, _ := db.Init()
	var contact *model.Contact
	contactCollection := client.Database("myutilityx").Collection("contacts")
	err := contactCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

// Helper function to create a pointer to a time.Time value
func timePtr(t time.Time) *time.Time {
	return &t
}
