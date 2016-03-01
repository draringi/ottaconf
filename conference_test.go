package ottaconf

import (
	"encoding/xml"
	"io/ioutil"
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
