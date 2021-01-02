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

type accountRepo struct {
	db *mongo.Client
}

func NewAccountRepo(db *mongo.Client) (*accountRepo, error) {

	return &accountRepo{
		db: db,
	}, nil
}

func (repo *accountRepo) Add(a model.Account) (err error) {
	collection := repo.db.Database("test").Collection("users")
	_, err = collection.InsertOne(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("New User created ")
	return
}

func (repo *accountRepo) Get(username string) (a model.Account, err error) {
	collection := repo.db.Database("test").Collection("users")
	filter := bson.D{{"username", username}}
	err = collection.FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", a)
	return
}

func (repo *accountRepo) Get1(id string) (a model.Account, err error) {
	collection := repo.db.Database("test").Collection("users")
	filter := bson.D{{"id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", a)
	return
}

func (repo *accountRepo) Update(userName string, name string) (err error) {
	collection := repo.db.Database("test").Collection("users")
	filter := bson.D{{"username", userName}}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
		}},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name user updated !")
	return
}

func (repo *accountRepo) GetAll() (allAccounts []model.Account, err error) {
	findOptions := options.Find()
	collection := repo.db.Database("test").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var a model.Account
		err := cur.Decode(&a)
		if err != nil {
			log.Fatal(err)
		}

		allAccounts = append(allAccounts, a)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return
}
