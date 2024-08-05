package main

import (
	"noox/cmd/api"
	"noox/db"
)

func main() {
	db.Ping()
	server := api.NewAPIServer(":8000")
	server.Run()
}
