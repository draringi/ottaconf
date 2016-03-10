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

func Diff(oldSchedule, newSchedule *ottaconf.Conference) []Change {
	changes := []Change{diffConferenceMeta(oldSchedule, newSchedule)}
	for _, e := range newSchedule.Events() {
		change := checkEvent(e, oldSchedule)
		if change != nil {
			changes = append(changes, change)
		}
	}
	// TODO: Add more diffing
	return changes
}
