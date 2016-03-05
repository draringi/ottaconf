package ottaconf

import "fmt"

// Time stores the times and durations found in the XML schedule in a machine friendly manner.
type Time int

// String returns the time in HH:MM format
func (t Time) String() string {
	hours := int(t) / 60
	mins := int(t) - hours*60
	return fmt.Sprintf("%0.2d:%0.2d", hours, mins)
}

// AtoTime converts a string in HH:MM format to the internal representation of time.
func AtoTime(str string) (Time, error) {
	var hours, mins int
	_, err := fmt.Sscanf(str, "%2d:%2d", &hours, &mins)
	if err != nil {
		return Time(0), err
	}
	return Time(hours*60 + mins), err
}
