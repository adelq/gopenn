package penn

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
