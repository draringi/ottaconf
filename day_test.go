package ottaconf

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestDayUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/day.xml")
	if err != nil {
		t.Fatalf("Error reading test data: %v", err)
	}
	d := new(Day)
	err = xml.Unmarshal(data, d)
	if err != nil {
		t.Errorf("Unable to Parse data: %v", err)
		return
	}
	if len(d.events) != 1 {
		t.Errorf("Day should have 1 Event, has %d", len(d.events))
	}
	e := d.events[0]
	if e.id != 605 {
		t.Errorf("Invalid event id. Expected 605, Found %d", e.id)
	}
	if e.title != "Goat BOF" {
		t.Errorf("Invalid event title. Expected \"Goat BOF\", Found %s", e.title)
	}
	if len(e.subtitle) != 0 {
		t.Errorf("Event subtitle should be empty. Has length %d", len(e.subtitle))
	}
	if len(e.people) != 3 {
		t.Errorf("Event should have 3 people assigned, has %d", len(e.people))
	}
	if len(e.links) != 2 {
		t.Errorf("Event should have 2 links, has %d", len(e.links))
	}
}
