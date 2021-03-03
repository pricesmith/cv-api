package congressApi

import (
	"net/http"
	"net/url"
	"time"
)

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

// not correct.. or necessary? 
func (c *Client) getMembers(queryParams) ([]Members, error) {
    req, err := http.NewRequest("GET", /*URL*/, nil)
    if err != nil {
        return nil, err
    }

    /*authenticate the request*/
    /*add queryParams to request*/

    res, err := c.c.Do()
    if err != nil {
        return nil, err
    }

    /*deserialize success or error response and return its data*/
}