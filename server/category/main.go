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
	"github.com/bhuvnesh13396/PlayMySong/category"
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
	songSvcAddr    = "http://localhost:8081"
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

	categoryRepo, err := mongoDB.NewCategoryRepo(db)
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

	categoryService := category.NewService(songService, categoryRepo)
	categoryService = category.NewLogService(categoryService, kit.LoggerWith(logger, "service", "Category"))
	categoryHandler := category.NewHandler(categoryService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)

	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	r.Handle("/song", songHandler)
	r.Handle("/song/", songHandler)

	r.Handle("/category", categoryHandler)
	r.Handle("/category/", categoryHandler)

	log.Println("listening on", ":8082")
	err = http.ListenAndServe(":8082", r)
	if err != nil {
		errExit(err)
	}
}
