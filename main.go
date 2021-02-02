package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/upload"

	"github.com/bhuvnesh13396/PlayMySong/account"

	"github.com/bhuvnesh13396/PlayMySong/auth"
	"github.com/bhuvnesh13396/PlayMySong/category"
	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	"github.com/bhuvnesh13396/PlayMySong/like"
	"github.com/bhuvnesh13396/PlayMySong/playlist"
	"github.com/bhuvnesh13396/PlayMySong/repo/local"
	mongoDB "github.com/bhuvnesh13396/PlayMySong/repo/mongo"
	"github.com/bhuvnesh13396/PlayMySong/song"
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

	sessionRepo, err := mongoDB.NewSessionRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	songRepo, err := mongoDB.NewSongRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	likeRepo, err := mongoDB.NewLikeRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	playlistRepo, err := mongoDB.NewPlaylistRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	categoryRepo, err := mongoDB.NewCategoryRepo(client)
	if err != nil {
		log.Fatal(err)
	}

	uploadRepo, err := local.NewSongUploadRepo()
	if err != nil {
		log.Fatal(err)
	}

	authService := auth.NewService(sessionRepo, accountRepo)
	authHandler := auth.NewHandler(authService)

	accountService := account.NewService(accountRepo)
	accountService = account.NewLogService(accountService, kit.LoggerWith(logger, "service", "Account"))
	// accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	songService := song.NewService(songRepo, accountRepo)
	songService = song.NewLogService(songService, kit.LoggerWith(logger, "service", "Song"))
	songHandler := song.NewHandler(songService)

	likeService := like.NewService(likeRepo)
	likeService = like.NewLogService(likeService, kit.LoggerWith(logger, "service", "Like"))
	likeHandler := like.NewHandler(likeService)

	playlistService := playlist.NewService(songRepo, playlistRepo)
	playlistService = playlist.NewLogService(playlistService, kit.LoggerWith(logger, "service", "Playlist"))
	playlistHandler := playlist.NewHandler(playlistService)

	categoryService := category.NewService(songRepo, categoryRepo)
	categoryService = category.NewLogService(categoryService, kit.LoggerWith(logger, "service", "Category"))
	categoryHandler := category.NewHandler(categoryService)

	uploadService := upload.NewService(uploadRepo)
	uploadService = upload.NewLogService(uploadService, kit.LoggerWith(logger, "service", "UploadService"))
	uploadHandler := upload.NewHandler(uploadService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)

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

	r.Handle("/upload", uploadHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}
