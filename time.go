package ottaconf

import "fmt"

type Time int

func (t Time) String() string {
	hours := int(t) / 60
	mins := int(t) - hours*60
	return fmt.Sprintf("%0.2d:%0.2d", hours, mins)
}

func AtoTime(str string) (Time, error) {
	var hours, mins int
	_, err := fmt.Sscanf(str, "%2d:%2d", &hours, &mins)
	if err != nil {
		return Time(0), err
	}
	return Time(hours*60 + mins), err
}
