package gopenn

import "net/http"

type DiningService interface {
	GetDaily(int) (*DiningHallDaily, *http.Response, error)
	GetWeekly(int) (*DiningHallWeekly, *http.Response, error)
	List() ([]Venue, *http.Response, error)
}

type DiningServiceOp struct {
	client *Client
}

// var _ DiningService = &DiningServiceOp{}

type DiningWrap struct {
	Venues []Venue     `json:"result_data.document.venue"`
	Meta   ServiceMeta `json:"service_meta"`
}

// Expected JSON structure from results from venues
type Venue struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	VenueType     string      `json:"venueType"`
	FacilityURL   string      `json:"facilityURL"`
	WeeklyMenuURL string      `json:"weeklyMenuURL"`
	DailyMenuURL  string      `json:"dailyMenuURL"`
	DateHours     []DateHours `json:"dateHours"`
}

type DateHours struct {
	Date  string `json:"date"`
	Meals []Meal `json:"meal"`
}

type Meal struct {
	OpenTime  string `json:"open"`
	CloseTime string `json:"close"`
	Type      string `json:"type"`
}

type DiningHallDaily struct {
}

type DiningHallWeekly struct {
}
