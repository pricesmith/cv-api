package main

import (
	"log"
	"os"

	"github.com/pricesmith/cv-api/clients"
	"github.com/pricesmith/cv-api/clients/propublica"
	"github.com/pricesmith/cv-api/server"
)

func main() {
	APIKey := os.Getenv("API_KEY")
	if APIKey == "" {
		log.Fatal("missing api key")
	}
	ppClient := propublica.New(APIKey)
	clients := clients.Clients{
		Propublica: *ppClient,
	}
	server := server.NewServer(clients)
	server.BuildRoutes()
	server.Run()
}
