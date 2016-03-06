package ottaconf

import (
	"errors"
	"fmt"
)

var datemap = make(map[string]*Date)

// Date stores the dates found in the XML schedule in a machine friendly manner.
type Date struct {
	year  uint16
	month uint8
	day   uint8
}

// String returns the date in YYYY-MM-DD format.
func (d *Date) String() string {
	return fmt.Sprintf("%0.4d-%0.2d-%0.2d", d.year, d.month, d.day)
}

// AtoDate converts a string in YYYY-MM-DD format to the internal representation of a date.
func AtoDate(str string) (*Date, error) {
	d, ok := datemap[str]
	if ok {
		return d, nil
	}
	d = new(Date)
	n, err := fmt.Sscanf(str, "%4d-%2d-%2d", &d.year, &d.month, &d.day)
	if err != nil {
		return nil, err
	} else if n != 3 {
		return nil, errors.New("Invalid date string")
	}
	datemap[str] = d
	return d, nil
}

// Year returns the year the date is in.
func (d *Date) Year() uint16 {
	return d.year
}

// Month returns the month the date is in, with 1 being January, and 12 being December.
func (d *Date) Month() uint8 {
	return d.month
}

// Day returns the day of the month the date is.
func (d *Date) Day() uint8 {
	return d.day
}
