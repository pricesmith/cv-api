package propublica

import (
	"net/http"
	"net/url"
	"time"
)

const PropublicaURL = "https://api.propublica.org/"

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.propublica.org",
	Path:   "/v1/",
}

type Client struct {
	c      *http.Client
	apiKey string
}

func New(apiKey string) *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}
