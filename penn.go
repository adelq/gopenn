package penn

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL string = "https://esb.isc-seo.upenn.edu/8091/open_data/"
	mediaType      string = "application/json; charset=utf-8"
)

// Client manages communiction with the Penn OpenData API.
type Client struct {
	client   *http.Client
	BaseURL  *url.URL
	username string
	password string
}

// NewClient returns a new Penn OpenData API client.
func NewClient(username string, password string) *Client {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{}
	c.client = httpClient
	c.BaseURL = baseURL
	c.username = username
	c.password = password
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// which will be resolved to the BaseURL of the Client. Relative URLs should
// always be specified without a preceding slash. If specified, the value
// pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("Authorization-Bearer", c.username)
	req.Header.Add("Authorization-Token", c.password)
	return req, nil
}
