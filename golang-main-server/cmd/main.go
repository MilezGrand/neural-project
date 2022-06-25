package main

import (
	"github.com/m4yb3/neural-project-server/internal/database"
	"github.com/m4yb3/neural-project-server/internal/server"
)

func main() {
	client, err := database.NewDatabase("root", "123456", "127.0.0.1", "hackhaton")
	if err != nil {
		panic(err)
	}

	err = server.NewServer(client)
	if err != nil {
		panic(err)
	}
}
