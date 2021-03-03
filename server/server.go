package server

import (
	"cv-api/congressApi"
	"net/http"
)

// os.Getenv("PPC_API_KEY"))?
var PPC_API_KEY string = "3vU0Ma4NFX2fLdqvNmats1qyTjASenhG5VrED6Gn"
var PROPUBLICA_URL string = "https://api.propublica.org/"

func Handlers() {
	http.HandleFunc("/home", congressApi.naHandler))
	http.HandleFunc("/blog", congressApi.naHandler))
	http.HandleFunc("/learn", congressApi.naHandler))
	http.HandleFunc("/forum", congressApi.naHandler))
	http.HandleFunc("/reads", congressApi.naHandler))
	http.HandleFunc("/about", congressApi.naHandler))
	http.HandleFunc("/contact", congressApi.naHandler))
	http.HandleFunc("/resources", congressApi.naHandler))

	http.HandleFunc("/search", congressApi.naHandler))
	http.HandleFunc("/search/persons", congressApi.naHandler))
	http.HandleFunc("/search/numbers", congressApi.naHandler))
	http.HandleFunc("/search/addresses", congressApi.naHandler))
	http.HandleFunc("/search/socialmedia", congressApi.naHandler))

	http.HandleFunc("/congress", congressApi.getMembersHandler)
	http.HandleFunc("/search/congress", congressApi.naHandler))
	http.HandleFunc("/search/corporate", congressApi.naHandler))
	http.HandleFunc("/search/businesses", congressApi.naHandler))

	http.HandleFunc("/polls", congressApi.naHandler))
	http.HandleFunc("/polls/create", congressApi.naHandler))
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
