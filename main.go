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

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

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

	entry := "pkg/http/web/app/dist/index.html"
	fs := http.FileServer(http.Dir("pkg/http/web/app/dist/assets/"))

	webService := web.New(repo, fs, entry)

	httpPort := host + ":" + port
	log.Printf("Running on host:port %s\n", httpPort)

	srv := &http.Server{
		Addr:         httpPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      webService.Router,
	}
	if port == "443" || port == "https" {
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))
		log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	} else {
		log.Fatal(srv.ListenAndServe())
	}

}
