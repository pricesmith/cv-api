package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pricesmith/cv-api/clients"
)

// Server contains server important information
type Server struct {
	host, port string
	Clients    clients.Clients
}

// NewServer creates a new server with client data
func NewServer(clients clients.Clients) *Server {
	return &Server{
		host:    "localhost",
		port:    "8080",
		Clients: clients,
	}
}

func (s *Server) BuildRoutes() {
	http.HandleFunc("/home", naHandler)
	http.HandleFunc("/blog", naHandler)
	http.HandleFunc("/learn", naHandler)
	http.HandleFunc("/forum", naHandler)
	http.HandleFunc("/reads", naHandler)
	http.HandleFunc("/about", naHandler)
	http.HandleFunc("/contact", naHandler)
	http.HandleFunc("/resources", naHandler)

	http.HandleFunc("/search", naHandler)
	http.HandleFunc("/search/persons", naHandler)
	http.HandleFunc("/search/numbers", naHandler)
	http.HandleFunc("/search/addresses", naHandler)
	http.HandleFunc("/search/socialmedia", naHandler)

	http.HandleFunc("/congress", s.getMembersHandler)
	http.HandleFunc("/search/congress", naHandler)
	http.HandleFunc("/search/corporate", naHandler)
	http.HandleFunc("/search/businesses", naHandler)

	http.HandleFunc("/polls", naHandler)
	http.HandleFunc("/polls/create", naHandler)
}

func (s *Server) Run() {
	fmt.Printf("Starting server at port %s\n", s.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil); err != nil {
		log.Fatal(err)
	}
}

func naHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route not available")
	return
}
