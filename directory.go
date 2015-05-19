package penn

import "net/http"

type DirectoryService interface {
	Get(string) (*Person, *http.Response, error)
	Search(*DirectorySearchOptions) ([]Person, *http.Response, error)
}

type DirectoryServiceOp struct {
	client *Client
}

// Expected JSON structure from results from directory
type Person struct {
	AdditionalData string `json:"additional_data_available"`
	Affiliation    string `json:"list_affiliation"`
	Email          string `json:"list_email"`
	Name           string `json:"list_name"`
	Organization   string `json:"list_organization"`
	Phone          string `json:"list_phone"`
	PhoneType      string `json:"list_phone_type"`
	Prefix         string `json:"list_prefix"`
	TitleOrMajor   string `json:"list_title_or_major"`
	Id             string `json:"person_id"`
}

// Acceptable parameters for use with the directory search endpoint.
type DirectorySearchOptions struct {
	FirstName    string `url:"first_name,omitempty"`
	Lastname     string `url:"last_name,omitempty"`
	Email        string `url:"email,omitempty"`
	Affiliation  string `url:"affiliation,omitempty"`
	Organization string `url:"organization,omitempty"`
}
