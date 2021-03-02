package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// os.Getenv("PPC_API_KEY"))?
var PPC_API_KEY string = "3vU0Ma4NFX2fLdqvNmats1qyTjASenhG5VrED6Gn"
var PROPUBLICA_URL string = "https://api.propublica.org/"

func Handlers() {
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

	http.HandleFunc("/congress", getMembersHandler)
	http.HandleFunc("/search/congress", naHandler)
	http.HandleFunc("/search/corporate", naHandler)
	http.HandleFunc("/search/businesses", naHandler)

	http.HandleFunc("/polls", naHandler)
	http.HandleFunc("/polls/create", naHandler)
}
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
	req.Header.Add("X-API-Key", PPC_API_KEY)

	// client := &http.Client????
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	var ppcRes PPCResponse
	if err := json.Unmarshal([]byte(responseBytes), &ppcRes); err != nil {
		log.Fatalf("error deserializing PPC response data")
	}
	log.Println(ppcRes)

	fmt.Fprintf(w, string(responseBytes))
}

// jsonString := `{
// 	"status": "OK",
// 	"copyright": "Copyright (c) 2021 Pro Publica Inc. All Rights Reserved.",
// 	"results": [
// 		{
// 			"congress": 	"116",
// 			"chamber": 		"Senate",
// 			"num_results": 	102,
// 			"offset":		0,
// 			"members":		[
// 				{
// 					"id": "A000360",
//             		"title": "Senator, 2nd Class",
// 					"short_title": "Sen.",
//             		"api_uri": "https://api.propublica.org/congress/v1/members/A000360.json",
//             		"first_name": "Lamar",
// 					"middle_name": null,
// 					"last_name": "Alexander",
//             		"suffix": null,
// 					"date_of_birth": "1940-07-03",
//             		"gender": "M",
//             		"party": "R",
// 					"leadership_role": null,
// 					"twitter_account": "SenAlexander",
//             		"facebook_account": "senatorlamaralexander",
//             		"youtube_account": "lamaralexander",
// 					"govtrack_id": "300002",
// 					"cspan_id": "5",
// 					"votesmart_id": "15691",
// 					"icpsr_id": "40304",
// 					"crp_id": "N00009888",
// 					"google_entity_id": "/m/01rbs3",
// 					"fec_candidate_id": "S2TN00058",
// 					"url": "https://www.alexander.senate.gov/public",
// 					"rss_url": "https://www.alexander.senate.gov/public/?a=RSS.Feed",
// 					"contact_form": "http://www.alexander.senate.gov/public/index.cfm?p=Email",
// 					"in_office": false,
// 					"cook_pvi": null,
// 					"dw_nominate": 0.324,
//             		"ideal_point": null,
// 					"cook_pvi": null,
// 					"seniority": "17",
// 					"next_election": "2020",
// 					"total_votes": 717,
// 					"missed_votes": 133,
// 					"total_present": 0,
// 					"last_updated": "2020-12-30 19:01:18 -0500",
// 					"ocd_id": "ocd-division/country:us/state:tn",
// 					"office": "455 Dirksen Senate Office Building",
// 					"phone": "202-224-4944",
// 					"fax": "202-228-3398",
// 					"state": "TN",
// 					"senate_class": "2",
// 					"state_rank": "senior",
// 					"lis_id": "S289"
// 					,"missed_votes_pct": 18.55,
// 					"votes_with_party_pct": 96.55,
// 					"votes_against_party_pct": 3.45
// 				}
// 			]
// 		}
// 	]
// }`
