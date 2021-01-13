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

type playlistRepo struct {
	db *mongo.Client
}

func NewPlaylistRepo(db *mongo.Client) (*playlistRepo, error) {
	return &playlistRepo{
		db: db,
	}, nil
}

func (repo *playlistRepo) Add(p model.Playlist) (err error) {
	collection := repo.db.Database("test").Collection("playlists")
	_, err = collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("New Playlist created ")
	return
}

func (repo *playlistRepo) Get(ID string) (p model.Playlist, err error) {
	collection := repo.db.Database("test").Collection("playlists")
	filter := bson.D{{"id", ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&p)
	fmt.Println("err", err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", p)
	return p, nil
}

func (repo *playlistRepo) Update(ID string, songIDs []string) (err error) {
	collection := repo.db.Database("test").Collection("playlists")
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

	fmt.Println("Playlists Updated !")
	return
}

func (repo *playlistRepo) List() (playlists []model.Playlist, err error) {
	findOptions := options.Find()
	collection := repo.db.Database("test").Collection("playlists")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var p model.Playlist
		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}

		playlists = append(playlists, p)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return
}
