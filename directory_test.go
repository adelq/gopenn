package gopenn

import "testing"

func TestDirectory_GetPerson(t *testing.T) {
	setup()
	defer teardown()

	person, _, err := client.Directory.Get("644cb460c2dc2d22046cd60e237ab534")
	if err != nil {
		t.Errorf("Directory.Get returned error: %v", err)
	}

	expected := "KNOP, SUZANNE E"
	if person.Name != expected {
		t.Errorf("Directory.Get returned %+v, expected %+v", person, expected)
	}
}

func TestDirectory_Search(t *testing.T) {
	setup()
	defer teardown()

	opt := &DirectorySearchOptions{LastName: "Knop"}
	personList, _, err := client.Directory.Search(opt)
	if err != nil {
		t.Errorf("Directory.Search returned error: %v", err)
	}

	expected := "KNOP, SUZANNE E"
	if personList[0].Name != expected {
		t.Errorf("Directory.Search returned %+v, expected %+v", personList, expected)
	}
}
