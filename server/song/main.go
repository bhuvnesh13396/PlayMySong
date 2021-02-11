package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/account"
	"github.com/bhuvnesh13396/PlayMySong/account/client"

	"github.com/bhuvnesh13396/PlayMySong/auth"
	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	mongoDB "github.com/bhuvnesh13396/PlayMySong/repo/mongo"
	"github.com/bhuvnesh13396/PlayMySong/song"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	accountSvcAddr = "http://localhost:8080"
)

func errExit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	db, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	logger := kit.NewJSONLogger(os.Stdout)

	accountRepo, err := mongoDB.NewAccountRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	sessionRepo, err := mongoDB.NewSessionRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	songRepo, err := mongoDB.NewSongRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	authService := auth.NewService(sessionRepo, accountRepo)
	authHandler := auth.NewHandler(authService)

	accountService, err := client.New(accountSvcAddr, http.DefaultClient)
	accountService = account.NewLogService(accountService, kit.LoggerWith(logger, "service", "Account"))
	// accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	songService := song.NewService(songRepo, accountService)
	songService = song.NewLogService(songService, kit.LoggerWith(logger, "service", "Song"))
	songHandler := song.NewHandler(songService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)

	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	r.Handle("/song", songHandler)
	r.Handle("/song/", songHandler)

	log.Println("listening on", ":8081")
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		errExit(err)
	}
}
