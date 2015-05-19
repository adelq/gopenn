package gopenn

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL string = "https://esb.isc-seo.upenn.edu/8091/open_data/"
	mediaType      string = "application/json; charset=utf-8"
)

// Client manages communiction with the Penn OpenData API.
type Client struct {
	client  *http.Client
	BaseURL *url.URL

	// Credentials for API
	username string
	password string

	// Services for communicating with API
	Directory DirectoryService
}

type ServiceMeta struct {
	CurrentPage        int    `json:"current_page_number"`
	ErrorText          string `json:"error_text"`
	NextPage           int    `json:"next_page_number"`
	NumberPages        int    `json:"number_of_pages"`
	PreviousPageNumber int    `json:"previous_page_number"`
	ResultsPerPage     int    `json:"results_per_page"`
}

// addOptions adds an interface of url options to a request
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qv, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qv.Encode()
	return u.String(), nil
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
	c.Directory = &DirectoryServiceOp{client: c}
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

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred, If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}
