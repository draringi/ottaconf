package ottaconf

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestPersonUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/person.xml")
	if err != nil {
		t.Fatalf("Unable to read test data: %v", err)
	}
	p := new(Person)
	err = xml.Unmarshal(data, &p)
	if err != nil {
		t.Errorf("Unable to Unmarshal: %v", err)
		return
	}
	if p.id != 214 {
		t.Errorf("Failed matching id. Expected 214, Found %d", p.id)
	}
	if p.name != "Allan Jude" {
		t.Errorf("Failed matching name. Expected \"Allan Jude\", Found \"%s\"", p.name)
	}
}
