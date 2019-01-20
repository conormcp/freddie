package main

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func main() {
	var err error
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Println("MONGO_URL not provided. Falling back to localhost:27017")
		mongoURL = "localhost:27017"
	}
	session, err := mgo.Dial(mongoURL)
	defer session.Close()

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("meetings").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
