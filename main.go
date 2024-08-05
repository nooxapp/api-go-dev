package main

import "noox/cmd/api"

func main() {
	server := api.NewAPIServer(":8000")
	server.Run()
}
