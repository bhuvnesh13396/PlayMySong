package mongo

import (
	"context"
	"log"

	"github.com/bhuvnesh13396/PlayMySong/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type sessionRepo struct {
	db *mongo.Client
}

func NewSessionRepo(db *mongo.Client) (*sessionRepo, error) {
	return &sessionRepo{
		db: db,
	}, nil
}

func (repo *sessionRepo) Add(s model.Session) (err error) {
	collection := repo.db.Database("test").Collection("sessions")
	_, err = collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func (repo *sessionRepo) Get(token string) (s model.Session, err error) {
	collection := repo.db.Database("test").Collection("sessions")
	filter := bson.D{{"token", token}}
	err = collection.FindOne(context.TODO(), filter).Decode(&s)

	if err != nil {
		log.Fatal(err)
	}

	return s, nil
}
