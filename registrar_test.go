package gopenn

import (
	"os"
	"testing"
)

func TestRegistrar_GetDepartment(t *testing.T) {
	setup()
	client.username = os.Getenv("REGISTRAR_API_USERNAME")
	client.password = os.Getenv("REGISTRAR_API_PASSWORD")
	defer teardown()

	department, _, err := client.Registrar.GetDepartment("CIS")
	if err != nil {
		t.Errorf("Registrar.GetDepartment returned error: %v", err)
	}

	expected := "CIS "
	if department[0].Department != expected {
		t.Errorf("Registrar.GetDepartment returned %+v, expected %+v", department[0].Department, expected)
	}
}

func TestRegistrar_GetCourse(t *testing.T) {
	setup()
	client.username = os.Getenv("REGISTRAR_API_USERNAME")
	client.password = os.Getenv("REGISTRAR_API_PASSWORD")
	defer teardown()

	course, _, err := client.Registrar.GetCourse("CIS", 110)
	if err != nil {
		t.Errorf("Registrar.GetCourse returned error: %v", err)
	}

	expected := "110"
	if course.CourseNumber != expected {
		t.Errorf("Registrar.GetCourse returned %+v, expected %+v", course.CourseNumber, expected)
	}
}

func TestRegistrar_Search(t *testing.T) {
	setup()
	client.username = os.Getenv("REGISTRAR_API_USERNAME")
	client.password = os.Getenv("REGISTRAR_API_PASSWORD")
	defer teardown()

	opt := &CourseSearchOptions{CourseID: "cis110"}
	courses, _, err := client.Registrar.Search(opt)
	if err != nil {
		t.Errorf("Registrar.Search returned error: %v", err)
	}

	expected := "110"
	if courses[0].CourseNumber != expected {
		t.Errorf("Registrar.Search returned %+v, expected %+v", courses[0].CourseID, expected)
	}
}
