package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/conormcp/freddie/pkg/ads"
	"github.com/conormcp/freddie/pkg/core"
	"github.com/globalsign/mgo/bson"
)

type Server struct {
	repo           core.Repository
	Router         *http.ServeMux
	entrypoint     string
	assets         http.Handler
	requestTimeout time.Duration
}

func (s Server) handleFrontend() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.entrypoint)
	})
}

func (s Server) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), s.requestTimeout)
}

func (s Server) handleAPI() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context for the request
		ctx, cancel := s.getContext()
		defer cancel()
		head, _ := core.ShiftPath(r.URL.Path)

		switch r.Method {
		case "GET":
			switch head {
			case "papers":
				ctx, _ = s.repo.NewContext(ctx, r)
				d := []core.Paper{}
				s.repo.FindAll(ctx, nil, &d, "-added")
				json.NewEncoder(w).Encode(d)
				return
			case "meetings":
				ctx, _ = s.repo.NewContext(ctx, r)
				d := []core.Meeting{}
				sel := bson.M{
					"date": bson.M{
						"$gte": time.Now(),
					},
				}
				s.repo.FindAll(ctx, sel, &d, "date")
				json.NewEncoder(w).Encode(d)
				return
			default:
				http.NotFound(w, r)
				return
			}
		case "POST":
			switch head {
			case "papers":
				ctx, _ = s.repo.NewContext(ctx, r)
				var d core.Paper
				json.NewDecoder(r.Body).Decode(&d)
				res, err := ads.Query(ctx, d.URL)
				if err != nil {
					fmt.Println(err)
				}
				d.Title = res[0].Title
				d.Abstract = res[0].Abstract
				d.Authors = res[0].Authors
				d.Added = time.Now()
				s.repo.Create(ctx, d)
				return
			case "meetings":
				ctx, _ = s.repo.NewContext(ctx, r)
				var d core.Meeting
				json.NewDecoder(r.Body).Decode(&d)
				fmt.Println(d)
				s.repo.Create(ctx, d)
				return
			default:
				http.NotFound(w, r)
				return
			}
		default:
			http.Error(w, "Method not allowd", http.StatusMethodNotAllowed)
			return
		}
	})
}

// New initializes a new Server
func New(repo core.Repository, fs http.Handler, entry string) Server {
	router := http.NewServeMux()

	s := Server{
		repo:           repo,
		Router:         router,
		entrypoint:     entry,
		assets:         fs,
		requestTimeout: time.Second * 60,
	}

	s.Router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	s.Router.HandleFunc("/api/", Chain(s.handleAPI(), Logging(), JSONApi(), Clip()))
	s.Router.HandleFunc("/", Chain(s.handleFrontend(), Logging(), Method("GET")))

	return s
}
