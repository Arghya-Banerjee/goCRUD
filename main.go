package main

import (
	"log"

	"github.com/Arghya-Banerjee/gocrud-api/database"
	"github.com/Arghya-Banerjee/gocrud-api/routes"
)

func main() {

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error Connecting to Database %v", err)
	}

	r := routes.SetupRouter()

	r.Run()

}