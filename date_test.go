package ottaconf

import "testing"

func TestAtoDate(t *testing.T) {
	str := "2016-02-29"
	date, err := AtoDate(str)
	if err != nil {
		t.Fatalf("Unable to parse \"%s\": %v", str, err)
	}
	if date.year != 2016 {
		t.Errorf("Wrong Year. Expected 2016, found %v", date.year)
	}
	if date.month != 2 {
		t.Errorf("Wrong Month. Expected 2, found %v", date.month)
	}
	if date.day != 29 {
		t.Errorf("Wrong Day. Expected 29, found %v", date.day)
	}
}

func TestDateString(t *testing.T) {
	str := "2016-02-29"
	date, err := AtoDate(str)
	if err != nil {
		t.Fatalf("Unable to parse \"%s\": %v", str, err)
	}
	if date.String() != str {
		t.Errorf("Wrong Date String. Expected \"%s\", found \"%s\"", str, date.String())
	}
}
