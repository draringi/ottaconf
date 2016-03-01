package ottaconf

import "testing"

func TestAtoTime(t *testing.T) {
	str := "04:04"
	time, err := AtoTime(str)
	if err != nil {
		t.Fatalf("Unable to parse \"%s\": %v", str, err)
	}
	if int(time) != 244 {
		t.Errorf("Incorrect Time Value. Expected 4*60+4 = 244, Found %d", int(time))
	}
}

func TestTimeString(t *testing.T) {
	str := "10:01"
	time, err := AtoTime(str)
	if err != nil {
		t.Fatalf("Unable to parse \"%s\": %v", str, err)
	}
	if time.String() != str {
		t.Errorf("Wrong Time String. Expected \"%s\", found \"%s\"", str, time.String())
	}
}
