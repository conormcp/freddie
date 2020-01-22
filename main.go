package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/conormcp/freddie/pkg/core"
	web "github.com/conormcp/freddie/pkg/http"
	"github.com/conormcp/freddie/pkg/storage"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
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

	godotenv.Load()
	addr := os.Getenv("SERVER_ADDR")


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

	m := &autocert.Manager{
		Cache:      autocert.DirCache("cache"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(addr),
        }

	srv := &http.Server{
		Addr:         httpPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      webService.Router,
		TLSConfig:    m.TLSConfig(),
	}

	if port == "443" || port == "https" {
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))
		log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	} else {
		log.Fatal(srv.ListenAndServe())
	}

}
