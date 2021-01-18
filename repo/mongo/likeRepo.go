package mongo

import (
	"fmt"
	"log"
	"context"
	"github.com/bhuvnesh13396/PlayMySong/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type likeRepo struct {
	db *mongo.Client
}

func NewLikeRepo(db *mongo.Client) (*likeRepo, error) {
	return &likeRepo{
		db: db,
	}, nil
}

func (repo *likeRepo) Add(a model.Like) (err error) {

	//Add entry in the likes table
	collection := repo.db.Database("test").Collection("likes")
	_, err = collection.InsertOne(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Increment the value of total likes for the given activity
	// in the aggregation table
	// aggreCollection := repo.db.Database("test").Collection("activities")
	// filter := bson.D{{"activity_id", activityID}, {"count": 1}}
	// var result bson.M
	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("Like added!")
	return
}

func (repo *likeRepo) Get(activityID string) (count int, err error) {
	collection := repo.db.Database("test").Collection("activities")
	filter := bson.D{{"activity_id", activityID}}
	var result bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result ", result)
	fmt.Printf("Found a single document: %+v\n", result["count"])

	return result["count"].(int), nil
}

func (repo *likeRepo) Delete(activityID string, userID string) (err error) {

	//Delete the entry from the likes table
	collection := repo.db.Database("test").Collection("likes")
	filter := bson.D{{"activity_id", activityID},{"user_id", userID}}
	_, err = collection.DeleteOne(context.TODO(), collection.FindOne(context.TODO(), filter))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Decrement the value of total likes for given activity
	// in the aggregation table
	//aggreCollection := repo.db.Database("test").Collection("activities")
	// _, err = collection.InsertOne(context.TODO(), a)
	fmt.Println("Like removed!")
	return
}

func (repo *likeRepo) UpdateCount(activityID string, count int) (err error) {
	collection := repo.db.Database("test").Collection("activities")
	filter := bson.D{{"activity_id", activityID}}
	update := bson.D{
		{"$set", bson.D{
			{"count", count},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Like count updated !")
	return
}
