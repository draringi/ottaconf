package ottadiff

import (
	"fmt"
	"github.com/draringi/ottaconf"
	"strconv"
)

type eventChange struct {
	changeType ChangeType
	changes    []string
	name       string
}

func (e *eventChange) Type() ChangeType {
	return c.changeType
}

func (e *eventChange) String() string {
	var result string
	switch e.eventType {
	case Modification:
		result = "Event " + c.name + " Changes:\n"
	case Insertation:
		result = "New Event " + c.name + ":\n"
	case Deletion:
		result = "Event " + c.name + "Deleted\n"
	default:
		return ""
	}
	for _, s := range c.changes {
		result += "\t" + s + "\n"
	}
	return result
}

func checkEvent(e2 *ottaconf.Event, c *ottaconf) *eventChange {
	e1, err := c.EventByID(e2.ID())
	if err == nil {
		return eventDiff(e1, e2)
	}
	e := new(EventChange)
	e.changeType = Insertation
	e.name = e2.Title()
	e.changes = append(e.changes, fmt.Sprintf("Title: %s", e2.Title()))
	e.changes = append(e.changes, fmt.Sprintf("Subtitle: %s", e2.Subtitle()))
	e.changes = append(e.changes, fmt.Sprintf("ID: %d", e2.ID()))
	e.changes = append(e.changes, fmt.Sprintf("Room: %v", e2.Room()))
	e.changes = append(e.changes, fmt.Sprintf("Date: %v", e2.Date()))
	e.changes = append(e.changes, fmt.Sprintf("Start Time: %v", e2.StartTime()))
	e.changes = append(e.changes, fmt.Sprintf("Duration: %v", e2.Duration()))
	e.changes = append(e.changes, fmt.Sprintf("Track: %v", e2.Track()))
	e.changes = append(e.changes, fmt.Sprintf("Type: %v", e2.Type()))
	return e
}
