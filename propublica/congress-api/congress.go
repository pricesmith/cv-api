package main

import "encoding/json"

// where do people overlap with members?
// perhaps create person object upon receiving this data?

// gotta create commented index of all this stuff.. mega documentation
// some fields need fixing
type Member struct {
	Id                   string      `json:"id"`
	Title                string      `json:"title"`
	TitleShort           string      `json:"short_title"`
	ApiUri               string      `json:"api_uri"`
	FirstName            string      `json:"first_name"`
	MiddleName           string      `json:"middle_name"`
	LastName             string      `json:"last_name"`
	Suffix               string      `json:"suffix"`
	DOB                  string      `json:"date_of_birth"`
	Gender               string      `json:"gender"`
	Party                string      `json:"party"`
	LeadershipRole       string      `json:"leadership_role"`
	Twitter              string      `json:"twitter_account"`
	Facebook             string      `json:"facebook_account"`
	Youtube              string      `json:"youtube_account"`
	GovtrackId           json.Number `json:"govtrack_id"`  // double check
	CspanId              json.Number `json:"cspan_id"`     // double check
	VotesmartId          json.Number `json:"votesmart_id"` // double check
	IcpsrId              json.Number `json:"icpsr_id"`     // what is this?
	CrpId                string      `json:"crp_id"`       // what is this?
	GoogleEntityId       string      `json:"google_entity_id"`
	FecCandidateId       string      `json:"fec_candidate_id"`
	Url                  string      `json:"url"` // rename to website? type of website??
	UrlRss               string      `json:"rss_url"`
	ContactForm          string      `json:"contact_form"` // eh.. details here?
	IsInOffice           *bool       `json:"in_office"`    // name ok?
	CookPvi              string      `json:"cook_pvi"`     // no idea what this is...
	DwNominate           float64     `json:"dw_nominate"`  // not *float64?? or just reg float?
	IdealPoint           string      `json:"ideal_point"`  // no idea
	Seniority            json.Number `json:"seniority"`
	NextElection         json.Number `json:"next_election"` // date type?? just receives a year I think..
	TotalVotes           int         `json:"total_votes"`   // not sure to what this is referring..
	MissedVotes          int         `json:"missed_votes"`  // maybe these are votes participated in or not?
	TotalPresent         int         `json:"total_present"` // nope
	LastUpdated          string      `json:"last_updated"`  // is this right? **** VERY IMPORTANT **** ex. "2020-12-30 19:01:18 -0500"
	OcdId                string      `json:"ocdd_id"`       // double check
	OfficeAddress        string      `json:"office"`        // maybe make address type
	Phone                string      `json:"phone"`         // obv needs to be converted
	Fax                  string      `json:"fax"`           // same goes here
	State                string      `json:"state"`         // only two letters is [2]bytes a thing?
	SenateClass          json.Number `json:"senate_class"`
	StateRank            string      `json:"state_rank"` // ex. "senior"
	LisId                string      `json:"lis_id"`
	MissedVotesPct       float64     `json:"missed_votes_pct"`     // what's pct?
	VotesWithPartyPct    float64     `json:"votes_with_party_pct"` // same
	VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
}

// https://stackoverflow.com/questions/49097385/how-to-decode-json-with-type-convert-from-string-to-integer-in-golang-for-non-lo
type Congress struct {
	Congress   json.Number `json:"congress"` // congressNo?
	Chamber    string      `json:"chamber"`
	NumResults int         `json:"num_results"`
	Offset     int         `json:"offset"`
	Members    []Member    `json:"members"`
}

// better name? PPResponseHead? idk
type PPCResponse struct {
	Status    string     `json:"status"`
	Copyright string     `json:"copyright"`
	Results   []Congress `json:"results"`
}
