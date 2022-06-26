package main

import (
	"github.com/thiago-henrique-leite/go-api/database"
	"github.com/thiago-henrique-leite/go-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
