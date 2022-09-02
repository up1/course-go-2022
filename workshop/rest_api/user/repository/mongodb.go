package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Firstname string             `bson:"first_name"`
	Lastname  string             `bson:"last_name"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
}

type mongoDb struct {
	client *mongo.Client
}

func NewMongoDb(mongo *mongo.Client) *mongoDb {
	return &mongoDb{
		client: mongo,
	}
}

func (repo *mongoDb) AddAccount(ctx context.Context, account *Account) (*Account, error) {
	collection := repo.client.Database("mydb").Collection("accounts")
	_, err := collection.InsertOne(context.Background(), account)
	if err != nil {
		logrus.Error("can't create account err: ", err)
		return nil, err
	}
	return account, nil
}
