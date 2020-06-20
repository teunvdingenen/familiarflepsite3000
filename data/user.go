package data

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	store *MongoDatastore
}

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.store.GetCollection("users").FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info(user)
	return &user, err
}

func (r *UserRepository) GetUser(ctx context.Context, args struct{ ID primitive.ObjectID }) (*User, error) {
	cur := r.store.GetCollection("users").FindOne(ctx, bson.D{{"_id", args.ID}})
	var user User
	err := cur.Decode(user)
	if err == nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetOneUser() (*User, error) {
	id, e := primitive.ObjectIDFromHex("5ee90bae5300f255a9ec1016")
	if e != nil {
		log.Error(e)
		return nil, e
	}
	var user User
	err := r.store.GetCollection("users").FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info(user)
	return &user, err
}

func (r *UserRepository) SaveUser(user User) (*primitive.ObjectID, error) {
	result, err := r.store.GetCollection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	objectId := result.InsertedID.(primitive.ObjectID)
	return &objectId, nil
}
