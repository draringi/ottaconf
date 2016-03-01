package ottaconf

import (
	"errors"
	"fmt"
)

var datemap = make(map[string]*Date)

type Date struct {
	year  uint16
	month uint8
	day   uint8
}

func (d Date) String() string {
	return fmt.Sprintf("%4d-%2d-%2d", d.year, d.month, d.day)
}

func AtoDate(str string) (*Date, error) {
	d, ok := datemap[str]
	if ok {
		return d, nil
	}
	d = new(Date)
	n, err := fmt.Sscanf(str, "%4d-%2d-%2d", &d.year, &d.month, &d.day)
	if err != nil {
		return nil, err
	} else if n != 1 {
		return nil, errors.New("Invalid date string")
	}
	datemap[str] = d
	return d, nil
}
