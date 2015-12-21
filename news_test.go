package gopenn

import (
	"os"
	"testing"
)

func TestNews_Get(t *testing.T) {
	setup()
	client.username = os.Getenv("NEWS_EVENTS_MAP_API_USERNAME")
	client.password = os.Getenv("NEWS_EVENTS_MAP_API_PASSWORD")
	defer teardown()

	news, _, err := client.News.Get("gutmann")
	if err != nil {
		t.Errorf("News.Get returned error: %v", err)
	}

	if len(news) < 1 {
		t.Errorf("News.Get returned %+v results, expected 1 or more", len(news))
	}
}
