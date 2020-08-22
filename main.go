package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bhuvnesh13396/PlayMySong/account"
	mongoDB "github.com/bhuvnesh13396/PlayMySong/repo/mongo"
	_ "github.com/lib/pq"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func errExit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// logger := kit.NewJSONLogger(os.Stdout)

	accountRepo, err := mongoDB.NewAccountRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	accountService := account.NewService(accountRepo)
	// accountService = account.NewLogService(accountService, kit.LoggerWith(logger, "service", "Account"))
	// accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	r := http.NewServeMux()

	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}
