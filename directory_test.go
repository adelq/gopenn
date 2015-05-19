package gopenn

import "testing"

func TestDirectory_GetPerson(t *testing.T) {
	setup()
	defer teardown()

	person, _, err := client.Directory.Get("4ad00e45edffd2ec2180673dabf4aace")
	if err != nil {
		t.Errorf("Directory.Get returned error: %v", err)
	}

	expected := "QALIEH, ADEL "
	if person.Name != expected {
		t.Errorf("Directory.Get returned %+v, expected %+v", person, expected)
	}
}

func TestDirectory_Search(t *testing.T) {
	setup()
	defer teardown()

	opt := &DirectorySearchOptions{LastName: "Wissmann"}
	personList, _, err := client.Directory.Search(opt)
	if err != nil {
		t.Errorf("Directory.Search returned error: %v", err)
	}

	expected := "WISSMANN, ALEXANDER R"
	if personList[0].Name != expected {
		t.Errorf("Directory.Search returned %+v, expected %+v", personList, expected)
	}
}
