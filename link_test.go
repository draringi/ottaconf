package ottaconf

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestLinkUnmarshal(t *testing.T) {
	testURL, testText := "https://twitter.com/dlangille/status/462732293686452224", "Twitter announcement"
	data, err := ioutil.ReadFile("testdata/link.xml")
	if err != nil {
		t.Fatalf("Unable to read test data: %v", err)
	}
	l := new(Link)
	err = xml.Unmarshal(data, &l)
	if err != nil {
		t.Errorf("Unable to Unmarshal: %v", err)
		return
	}
	if l.url != testURL {
		t.Errorf("Failed matching id.\n\tExpected \"%s\"\n\tFound \"%s\"", testURL, l.url)
	}
	if l.text != testText {
		t.Errorf("Failed matching name.\n\tExpected \"%s\"\n\tFound \"%s\"", testText, l.text)
	}
}
