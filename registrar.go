package gopenn

import (
	"fmt"
	"net/http"
)

type RegistrarService interface {
	GetDepartment(string) ([]Course, *http.Response, error)
	GetCourse(string, int) (*Course, *http.Response, error)
	Search(*CourseSearchOptions) ([]Course, *http.Response, error)
}

type RegistrarServiceOp struct {
	client *Client
}

var _ RegistrarService = &RegistrarServiceOp{}

// Expected JSON structure for course from course catalog
type Course struct {
	ActivitiesAndCredits []struct {
		Credit       string `json:"credit"`
		ActivityCode string `json:"activity_code"`
	} `json:"activities_and_credits"`
	Corequisites           string `json:"corequisites"`
	CourseCreditConnector  string `json:"course_credit_connector"`
	CourseCreditType       string `json:"course_credit_type"`
	CourseDescription      string `json:"course_description"`
	CourseID               string `json:"course_id"`
	CourseLevel            string `json:"course_level"`
	CourseLevelDescription string `json:"course_level_description"`
	CourseNotes            string `json:"course_notes"`
	CourseNumber           string `json:"course_number"`
	CourseTitle            string `json:"course_title"`
	Crosslistings          []struct {
		CourseID  string `json:"course_id"`
		SectioNID string `json:"section_id"`
		Subject   string `json:"subject"`
	} `json:"crosslistings"`
	Department              string `json:"department"`
	DepartmentOfRecord      string `json:"department_of_record"`
	DistributionRequirement string `json:"distribution_requirement"`
	EasCreditFactorCode     string `json:"eas_credit_factor_code"`
	// Instructors             []interface{} `json:"instructors"`
	Prerequisites           string        `json:"prerequisites"`
	RegisterSubgroupOne     string        `json:"register_subgroup_one"`
	RegisterSubgroupTwo     string        `json:"register_subgroup_two"`
	RequirementsMet         []interface{} `json:"requirements_met"`
	SchedulingPriority      string        `json:"scheduling_priority"`
	SchoolCode              string        `json:"school_code"`
	TermsOfferedCode        string        `json:"terms_offered_code"`
	TermsOfferedDescription string        `json:"terms_offered_description"`
}

type CourseWrap struct {
	Courses []Course    `json:"result_data"`
	Meta    ServiceMeta `json:"service_meta"`
}

// Acceptable parameters for use with the course search endpoint.
// Still missing many values due to lack of enums
type CourseSearchOptions struct {
	CourseID    string `url:"course_id,omitempty"`
	Description string `url:"description,omitempty"`
	Instructor  string `url:"instructor,omitempty"`
}

func (s *RegistrarServiceOp) GetDepartment(department string) ([]Course, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/", "course_info", department)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	courses := new(CourseWrap)
	resp, err := s.client.Do(req, courses)
	if err != nil {
		return nil, resp, err
	}

	return courses.Courses, resp, err
}

func (s *RegistrarServiceOp) GetCourse(department string, courseID int) (*Course, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/%d/", "course_info", department, courseID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	courses := new(CourseWrap)
	resp, err := s.client.Do(req, courses)
	if err != nil {
		return nil, resp, err
	}

	return &courses.Courses[0], resp, err
}

func (s *RegistrarServiceOp) Search(opt *CourseSearchOptions) ([]Course, *http.Response, error) {
	path := "course_section_search"
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	courses := new(CourseWrap)
	resp, err := s.client.Do(req, courses)
	if err != nil {
		return nil, resp, err
	}

	return courses.Courses, resp, err
}
