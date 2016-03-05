package ottaconf

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestConferenceUnmarshall(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/conference.xml")
	if err != nil {
		t.Fatalf("Unable to read test data: %v", err)
	}
	c := new(Conference)
	err = xml.Unmarshal(data, c)
	if err != nil {
		t.Errorf("Unable to Unmarshal: %v", err)
		return
	}
	if c == nil {
		t.Error("No Conference found")
		return
	}
	if c.title != "Meta-Conference" {
		t.Errorf("Failed matching title. Expected \"Meta-Conference\", found \"%s\"", c.title)
		return
	}
}

func TestParse(t *testing.T) {
	f, err := os.Open("testdata/bsdcan2015-schedule.en.xml")
	if err != nil {
		t.Fatalf("Unable to read test data: %v", err)
	}
	defer f.Close()
	c, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse conference: %v", err)
	}
	if c == nil {
		t.Error("No Conference found")
		return
	}
	if c.title != "BSDCan 2015" {
		t.Errorf("Conference title failed to match.\n\tExpected: \"BSDCan 2015\"\n\tFound: \"%s\"", c.title)
	}
	if c.length != 6 {
		t.Errorf("Conference should last 6 days, Found %d days", c.length)
	}
	if c.start.String() != "2015-06-09" {
		t.Errorf("Start date of conference fails to match.\n\tExpected 2015-06-09.\n\tFound %s.", c.start.String())
	}
	if c.end.String() != "2015-06-14" {
		t.Errorf("End date of conference fails to match.\n\tExpected 2015-06-14.\n\tFound %s.", c.end.String())
	}
	if len(c.rooms) != 14 {
		t.Errorf("Room count for conference does not match.\n\tExpected 15\n\tFound %d", len(c.rooms))
	}
	if len(c.days) != c.length {
		t.Errorf("Conference runs for %d days, but only %d days found", c.length, len(c.days))
	}
}
