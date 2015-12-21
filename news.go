package gopenn

import (
	"net/http"
	"time"
)

type NewsService interface {
	Get(string) ([]News, *http.Response, error)
}

type NewsServiceOp struct {
	client *Client
}

var _ NewsService = &NewsServiceOp{}

// Expected JSON structure for news articles
type News struct {
	Author            string `json:"author"`
	CommentsURL       string `json:"comments_url"`
	ContentCategories []struct {
		Description      string `json:"description"`
		ID               int    `json:"id"`
		Label            string `json:"label"`
		ParentCategoryID int    `json:"parent_category_id"`
	} `json:"content_categories"`
	ContentType string `json:"content_type"`
	Description string `json:"description"`
	GUID        string `json:"guid"`
	ID          int    `json:"id"`
	Images      []struct {
		Caption string `json:"caption"`
		ID      int    `json:"id"`
		Link    string `json:"link"`
		Name    string `json:"name"`
	} `json:"images"`
	Link      string    `json:"link"`
	Pubdate   time.Time `json:"pubdate"`
	SourceURL string    `json:"source_url"`
	Teaser    string    `json:"teaser"`
	Title     string    `json:"title"`
}

type NewsWrap struct {
	News []News      `json:"result_data"`
	Meta ServiceMeta `json:"service_meta"`
}

// Acceptable parameters for use with the news search endpoint.
type NewsSearchOptions struct {
	Source      string `url:"source,omitempty"`
	Description string `url:"description,omitempty"`
}

func (s *NewsServiceOp) Get(description string) ([]News, *http.Response, error) {
	path := "news_events_maps"
	opt := &NewsSearchOptions{Source: "news", Description: description}
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	news := new(NewsWrap)
	resp, err := s.client.Do(req, news)
	if err != nil {
		return nil, resp, err
	}

	return news.News, resp, err
}
