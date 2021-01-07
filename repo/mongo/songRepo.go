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

type songRepo struct {
	db *mongo.Client
}

func NewSongRepo(db *mongo.Client) (*songRepo, error) {

	return &songRepo{
		db: db,
	}, nil
}

func (repo *songRepo) Add(s model.Song) (err error) {
	collection := repo.db.Database("test").Collection("songs")
	_, err = collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("New song added to DB.")
	return
}

func (repo *songRepo) Get(songName string) (s model.Song, err error) {
	collection := repo.db.Database("test").Collection("songs")
	filter := bson.D{{"title", songName}}
	err = collection.FindOne(context.TODO(), filter).Decode(&s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", s)
	return
}

func (repo *songRepo) Get1(ID string) (s model.Song, err error) {
	collection := repo.db.Database("test").Collection("songs")
	filter := bson.D{{"id", ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&s)
	fmt.Println("err", err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", s)
	return s, nil
}

func (repo *songRepo) Update(id string, title string) (err error) {
	collection := repo.db.Database("test").Collection("songs")
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"title", title},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Song title updated !")
	return
}

func (repo *songRepo) List() (allSongs []model.Song, err error) {
	findOptions := options.Find()
	collection := repo.db.Database("test").Collection("songs")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var a model.Song
		err := cur.Decode(&a)
		if err != nil {
			log.Fatal(err)
		}

		allSongs = append(allSongs, a)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return
}
