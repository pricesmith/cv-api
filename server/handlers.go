package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getMembersHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["congress"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'congress' is missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	key := keys[0]
	results := s.Clients.Propublica.GetMembers(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
