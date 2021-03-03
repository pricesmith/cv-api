package congressApi

import (
	"cv-api/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func naHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route not available")
	return
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["congress"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'congress' is missing")
		return
	}
	key := keys[0]

	// url := "https://api.propublica.org/congress/v1/{congress}/{chamber}/members.json"
	url := "https://api.propublica.org/congress/v1/" + key + "/senate/members.json"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", server.PPC_API_KEY)

	// client := &http.Client????
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	// responseBytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalf("error reading HTTP response body: %v", err)
	// }

	var ppcRes PPCResponse
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&ppcRes); err != nil {
		log.Fatalf("error deserializing PPC response data")
	}
	log.Println(ppcRes)
}
