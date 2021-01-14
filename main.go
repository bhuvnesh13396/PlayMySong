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
	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	"github.com/bhuvnesh13396/PlayMySong/like"
	"github.com/bhuvnesh13396/PlayMySong/category"
	"github.com/bhuvnesh13396/PlayMySong/playlist"
	mongoDB "github.com/bhuvnesh13396/PlayMySong/repo/mongo"
	"github.com/bhuvnesh13396/PlayMySong/song"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	logger := kit.NewJSONLogger(os.Stdout)

	accountRepo, err := mongoDB.NewAccountRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	accountService := account.NewService(accountRepo)
	accountService = account.NewLogService(accountService, kit.LoggerWith(logger, "service", "Account"))
	// accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	songRepo, err := mongoDB.NewSongRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	songService := song.NewService(songRepo, accountRepo)
	songService = song.NewLogService(songService, kit.LoggerWith(logger, "service", "Song"))
	songHandler := song.NewHandler(songService)

	likeRepo, err := mongoDB.NewLikeRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	likeService := like.NewService(likeRepo)
	likeService = like.NewLogService(likeService, kit.LoggerWith(logger, "service", "Like"))
	likeHandler := like.NewHandler(likeService)

	playlistRepo, err := mongoDB.NewPlaylistRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	playlistService := playlist.NewService(songRepo, playlistRepo)
	playlistService = playlist.NewLogService(playlistService, kit.LoggerWith(logger, "service", "Playlist"))
	playlistHandler := playlist.NewHandler(playlistService)

	categoryRepo, err := mongoDB.NewCategoryRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	categoryService := category.NewService(songRepo, categoryRepo)
	categoryService = category.NewLogService(categoryService, kit.LoggerWith(logger, "service", "Category"))
	categoryHandler := category.NewHandler(categoryService)


	r := http.NewServeMux()

	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	r.Handle("/song", songHandler)
	r.Handle("/song/", songHandler)

	r.Handle("/like", likeHandler)
	r.Handle("/like/", likeHandler)

	r.Handle("/playlist", playlistHandler)
	r.Handle("/playlist/", playlistHandler)

	r.Handle("/category", categoryHandler)
	r.Handle("/category/", categoryHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}
