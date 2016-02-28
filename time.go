package ottaconf

type Time int

func (t Time) String() string {
	hours = int(t) / 60
	mins = int(t) - hours*60
	return fmt.Sprintf("%2d:%2d", hours, mins)
}

func AtoTime(str string) (Time, err) {
	var hours, mins int
	n, err := fmt.Sscanf(str, "%2d:%2d", &hours, &mins)
	if err != nil {
		return Time(0), err
	}
	return Time(hours*60 + mins), err
}
