package ottaconf

import (
	"io/ioutil"
	"testing"
	"encoding/xml"
)

func TestEventUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/event.xml")
	if err != nil {
		t.Fatalf("Error reading test data: %v", err)
	}
	e := new(Event)
	err = xml.Unmarshal(data, e)
	if err != nil {
		t.Errorf("Unable to Parse data: %v", err)
		return
	}
	if e.id != 605 {
		t.Errorf("Invalid event id. Expected 605, Found %d", e.id)
	}
	if e.title != "Goat BOF" {
		t.Errorf("Invalid event title. Expected \"Goat BOF\", Found %s", e.title)
	}
}
