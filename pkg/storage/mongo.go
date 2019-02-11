package storage

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/conormcp/freddie/pkg/core"
	"github.com/globalsign/mgo"
)

type key string

const (
	dbKey         = key("DB")
	collectionKey = key("table")
)

type MongoRepository struct {
	logger  *log.Logger
	session *mgo.Session
	dbName  string
}

func (r MongoRepository) fromContext(ctx context.Context) (*mgo.Session, *mgo.Collection) {
	session := r.session.Copy()
	collection, ok := ctx.Value(collectionKey).(string)
	if !ok {
		r.logger.Fatal("Failed to read context")
	}
	coll := session.DB(r.dbName).C(collection)
	return session, coll
}

// NewContext reads the context from the request and returns an updated context
func (r MongoRepository) NewContext(ctx context.Context, req *http.Request) (context.Context, error) {
	head, _ := core.ShiftPath(req.URL.Path)
	ctx = context.WithValue(ctx, collectionKey, head)
	return ctx, nil
}

// Find fetches a meeting from mongo according to the query criteria provided.
func (r MongoRepository) Find(ctx context.Context, id core.ID, result interface{}) error {
	session, coll := r.fromContext(ctx)
	defer session.Close()

	err := coll.FindId(id).One(result)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return err
	}
	return nil
}

// FindAll fetches all meetings from the database. YES.. ALL! be careful.
func (r MongoRepository) FindAll(ctx context.Context, selector interface{}, result interface{}, sort string) error {
	session, coll := r.fromContext(ctx)
	defer session.Close()

	err := coll.Find(selector).Sort(sort).All(result)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return err
	}
	return nil
}

// Delete deletes a meeting from mongo according to the query criteria provided.
func (r MongoRepository) Delete(ctx context.Context, id core.ID) error {
	session, coll := r.fromContext(ctx)
	defer session.Close()

	return coll.RemoveId(id)
}

// Update updates an meeting.
func (r MongoRepository) Update(ctx context.Context, id core.ID, item interface{}) error {
	session, coll := r.fromContext(ctx)
	defer session.Close()

	return coll.UpdateId(id, item)
}

// Create meetings in the database.
func (r MongoRepository) Create(ctx context.Context, items ...interface{}) error {
	session, coll := r.fromContext(ctx)
	defer session.Close()

	for _, i := range items {
		err := coll.Insert(i)
		if err != nil {
			return err
		}
	}

	return nil
}

// Count counts documents for a given collection
func (r MongoRepository) Count(ctx context.Context) (int, error) {
	session, coll := r.fromContext(ctx)
	defer session.Close()
	return coll.Count()
}

/*
func (r MongoRepository) Test(ctx context.Context) {
	m := &core.Meeting{Date: time.Now(), Title: "That Stuff", Speaker: "That guy"}
	fmt.Println(m)
	r.Create(ctx, m)
	p := &core.Paper{Added: time.Now(), Title: "A Treatise on That Stuff", Authors: "That guy et al.", URL: "http://www.google.com"}
	fmt.Println(p)
	r.SetTable("papers")
	r.Create(ctx, p)
	p = &core.Paper{
		Added:    time.Now(),
		URL:      "http://www.google.com",
		Title:    "Google",
		Authors:  "gophers",
		Abstract: "I am goolgebot!",
	}
	fmt.Println(p)
	r.Create(ctx, p)

}
*/

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB] ", 0)
}

func NewMongoRepository(dbName string) *MongoRepository {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL not provided")
	}
	logger := newMongoRepositoryLogger()

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v\n", err)
	}

	return &MongoRepository{
		session: session,
		logger:  logger,
		dbName:  dbName,
	}
}
