package main

import (
	"fmt"
	"os"

	"github.com/GeaRvant/go_finalProject/pkg/api"
	"github.com/GeaRvant/go_finalProject/pkg/db"
	"github.com/GeaRvant/go_finalProject/pkg/server"
)

func main() {
	webDir := "./web"
	port := ":7540"

	err := db.Init("scheduler.db")
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		if err := db.CloseDB(); err != nil {
			fmt.Println("Error closing database:", err)
		}
	}()

	api.Init()

	server.StartServer(port, webDir)
}
