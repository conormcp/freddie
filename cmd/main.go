package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/conormcp/freddie/pkg/core"
	web "github.com/conormcp/freddie/pkg/http"
	"github.com/conormcp/freddie/pkg/storage"
)

func main() {
	var host string
	var port string
	var dbURL string

	flag.StringVar(&host, "host", "localhost", "the `host` to listen on.")
	flag.StringVar(&port, "port", "3000", "the `port` to listen on.")
	flag.StringVar(&dbURL, "db-url", "mongodb://localhost:27017", "database url.")
	flag.Parse()

	var repo core.Repository
	repo = storage.NewMongoRepository("test")
	//repo.Test(context.Background())

	entry := "pkg/http/web/app/dist/index.html"
	fs := http.FileServer(http.Dir("pkg/http/web/app/dist/assets/"))

	webService := web.New(repo, fs, entry)

	httpPort := host + ":" + port
	log.Printf("Running on port %s\n", httpPort)

	srv := &http.Server{
		Addr:         httpPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      webService.Router,
	}
	log.Fatal(srv.ListenAndServe())
}
