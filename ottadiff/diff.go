package ottadiff

import "github.com/draringi/ottaconf"

type ChangeType int

const (
	Insertation ChangeType = iota
	Modification
	Deletion
)

type Change interface {
	Type() ChangeType
	String() string
}

func Diff(oldSchedule, newSchedule *ottaconf.Conference) []*Change {

}
