package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/bhuvnesh13396/PlayMySong/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type categoryRepo struct {
	db *mongo.Client
}

func NewCategoryRepo(db *mongo.Client) (*categoryRepo, error) {
	return &categoryRepo{
		db: db,
	}, nil
}

func (repo *categoryRepo) Add(c model.Category) (err error) {
	collection := repo.db.Database("test").Collection("categories")
	_, err = collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (repo *categoryRepo) Get(ID string) (c model.Category, err error) {
	collection := repo.db.Database("test").Collection("categories")
	filter := bson.D{{"id", ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)

	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}

func (repo *categoryRepo) Update(ID string, songIDs []string) (err error) {
	collection := repo.db.Database("test").Collection("categories")
	filter := bson.D{{"id", ID}}
	update := bson.D{
		{"$set", bson.D{
			{"song_ids", songIDs},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	
	return
}

func (repo *categoryRepo) List() (categories []model.Category, err error) {
	findOptions := options.Find()
	collection := repo.db.Database("test").Collection("categories")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var c model.Category
		err := cur.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, c)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return
}
