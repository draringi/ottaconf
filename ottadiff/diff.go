package ottadiff

import "github.com/draringi/ottaconf"

// ChangeType is an enumerator storing what type of change has taken place.
type ChangeType int

const (
	// Insertation signals an addition of an object to the schedule
	Insertation ChangeType = iota
	// Modification signals that an object in the schedule has been modified
	Modification
	// Deletion signals that an object has been removed from the schedule
	Deletion
)

// Change represents the any change within the schedule, to any object type.
type Change interface {
	Type() ChangeType
	String() string
}

// Diff examines two schedules, and returns a list of changes from the old schedule to the new schedule.
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
