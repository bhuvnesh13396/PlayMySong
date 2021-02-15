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
	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	"github.com/bhuvnesh13396/PlayMySong/like"
	"github.com/bhuvnesh13396/PlayMySong/playlist"
	"github.com/bhuvnesh13396/PlayMySong/repo/local"
	mongoDB "github.com/bhuvnesh13396/PlayMySong/repo/mongo"
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

	// var discoveryURL = flag.String("discovery", "127.0.0.1:8500", "Consul service discovery url")
	// var httpPort = flag.String("http", ":3000", "Port to run HTTP service at")
	//
	// flag.Parse()
	//
	// reg, err := registry.New(*discoveryURL)
	// if err != nil {
	// 	log.Fatalf("an error occurred while bootstrapping service discovery... %v", err)
	// }
	//
	// var healthURL string
	//
	// ip, err := registry.IPAddr()
	// if err != nil {
	// 	log.Fatalf("could not determine IP address to register this service with... %v", err)
	// }
	//
	// healthURL = "http://" + ip.String() + *httpPort + "/health"
	//
	// pp, err := strconv.Atoi((*httpPort)[1:]) // get rid of the ":" port
	// if err != nil {
	// 	log.Fatalf("could not discover port to register with consul.. %v", err)
	// }
	//
	// svc := &api.AgentServiceRegistration{
	// 	Name:    "cool_app",
	// 	Address: ip.String(),
	// 	Port:    pp,
	// 	Tags:    []string{"urlprefix-/oops"},
	// 	Check: &api.AgentServiceCheck{
	// 		TLSSkipVerify: true,
	// 		Method:        "GET",
	// 		Timeout:       "20s",
	// 		Interval:      "1m",
	// 		HTTP:          healthURL,
	// 		Name:          "HTTP check for cool app",
	// 	},
	// }
	//
	// id, err := reg.RegisterService(svc)
	// if err != nil {
	// 	log.Fatalf("Could not register service in consul... %v", err)
	// }
	//
	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	r.Body.Close()
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("OK"))
	// })
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	r.Body.Close()
	// 	fmt.Println("Here")
	// 	w.Write([]byte("home page"))
	// })
	//
	// if err := http.ListenAndServe(*httpPort, nil); err != nil {
	// 	reg.DeRegister(id)
	// }

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

	uploadRepo, err := local.NewSongUploadRepo()
	if err != nil {
		log.Fatal(err)
	}
	// uploadPath := "/home/bhuvi/work/src/github.com/bhuvnesh13396/PlayMySong/song_data/raw/123456789"
	// fs := http.FileServer(http.Dir(uploadPath))
	// http.Handle("/files/", http.StripPrefix("/files", fs))

	authService := auth.NewService(sessionRepo, accountRepo)
	authHandler := auth.NewHandler(authService)

	accountService := account.NewService(accountRepo)
	accountService = account.NewLogService(accountService, kit.LoggerWith(logger, "service", "Account"))
	// accountService = account.NewAuthService(accountService, authService)
	accountHandler := account.NewHandler(accountService)

	//songService := song.NewService(songRepo, accountRepo)
	//songService = song.NewLogService(songService, kit.LoggerWith(logger, "service", "Song"))
	//songHandler := song.NewHandler(songService)

	likeService := like.NewService(likeRepo)
	likeService = like.NewLogService(likeService, kit.LoggerWith(logger, "service", "Like"))
	likeHandler := like.NewHandler(likeService)

	playlistService := playlist.NewService(songRepo, playlistRepo)
	playlistService = playlist.NewLogService(playlistService, kit.LoggerWith(logger, "service", "Playlist"))
	playlistHandler := playlist.NewHandler(playlistService)

	uploadService := upload.NewService(uploadRepo)
	uploadService = upload.NewLogService(uploadService, kit.LoggerWith(logger, "service", "UploadService"))
	uploadHandler := upload.NewHandler(uploadService)

	r := http.NewServeMux()

	r.Handle("/auth", authHandler)
	r.Handle("/auth/", authHandler)

	r.Handle("/account", accountHandler)
	r.Handle("/account/", accountHandler)

	//r.Handle("/song", songHandler)
	//r.Handle("/song/", songHandler)

	r.Handle("/like", likeHandler)
	r.Handle("/like/", likeHandler)

	r.Handle("/playlist", playlistHandler)
	r.Handle("/playlist/", playlistHandler)

	r.Handle("/upload", uploadHandler)

	log.Println("listening on", ":8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		errExit(err)
	}
}
