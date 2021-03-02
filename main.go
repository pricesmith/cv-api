package main

import (
	"fmt"
	"log"
	"net/http"

	"server"
)

func main() {
	server.Handlers()
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
