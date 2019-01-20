#! /bin/sh
sudo GOPATH=/home/conor/gocode MONGO_URL=mongodb://localhost:27017  go run cmd/main.go -host= -port=https
