package gopenn

import (
	"fmt"
	"net/http"
)

type DirectoryService interface {
	Get(string) (*Person, *http.Response, error)
	Search(*DirectorySearchOptions) ([]Person, *http.Response, error)
}

type DirectoryServiceOp struct {
	client *Client
}

var _ DirectoryService = &DirectoryServiceOp{}

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

type PersonWrap struct {
	Persons []Person    `json:"result_data"`
	Meta    ServiceMeta `json:"service_meta"`
}

// Acceptable parameters for use with the directory search endpoint.
type DirectorySearchOptions struct {
	FirstName    string `url:"first_name,omitempty"`
	LastName     string `url:"last_name,omitempty"`
	Email        string `url:"email,omitempty"`
	Affiliation  string `url:"affiliation,omitempty"`
	Organization string `url:"organization,omitempty"`
}

func (s *DirectoryServiceOp) Get(personId string) (*Person, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", "directory_person_details", personId)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	person := new(PersonWrap)
	resp, err := s.client.Do(req, person)
	if err != nil {
		return nil, resp, err
	}

	return &person.Persons[0], resp, err
}

func (s *DirectoryServiceOp) Search(opt *DirectorySearchOptions) ([]Person, *http.Response, error) {
	path := "directory"
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	personList := new(PersonWrap)
	resp, err := s.client.Do(req, personList)
	if err != nil {
		return nil, resp, err
	}

	return personList.Persons, resp, err
}
