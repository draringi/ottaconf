package ottadiff

import (
	"fmt"
	"github.com/draringi/ottaconf"
)

type eventChange struct {
	changeType ChangeType
	changes    []string
	name       string
}

func (e *eventChange) Type() ChangeType {
	return e.changeType
}

func (e *eventChange) String() string {
	var result string
	switch e.changeType {
	case Modification:
		result = fmt.Sprintf("Event %s Changes:\n", e.name)
	case Insertation:
		result = fmt.Sprintf("New Event %s:\n", e.name)
	case Deletion:
		result = fmt.Sprintf("Event %s Deleted\n", e.name)
	default:
		return ""
	}
	for _, s := range e.changes {
		result += fmt.Sprintf("\t%s\n", s)
	}
	return result
}

func checkEvent(e2 *ottaconf.Event, c *ottaconf.Conference) *eventChange {
	e1, err := c.EventByID(e2.ID())
	if err == nil {
		return eventDiff(e1, e2)
	}
	e := new(eventChange)
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

func eventDiff(e1, e2 *ottaconf.Event) *eventChange {
	e := new(eventChange)
	e.changeType = Modification
	if e1.Title() != e2.Title() {
		e.changes = append(e.changes, fmt.Sprintf("Title: %v -> %v", e1.Title(), e2.Title()))
	}
	if e1.Subtitle() != e2.Title() {
		e.changes = append(e.changes, fmt.Sprintf("Subtitle: %v -> %v", e1.Subtitle(), e2.Subtitle()))
	}
	if e1.Room() != e2.Room() {
		e.changes = append(e.changes, fmt.Sprintf("Room: %v -> %v", e1.Room(), e2.Room()))
	}
	if e1.Date() != e2.Date() {
		e.changes = append(e.changes, fmt.Sprintf("Date: %v -> %v", e1.Date(), e2.Date()))
	}
	if e1.StartTime() != e2.StartTime() {
		e.changes = append(e.changes, fmt.Sprintf("Start Time: %v -> %v", e1.StartTime(), e2.StartTime()))
	}
	if e1.Duration() != e2.Duration() {
		e.changes = append(e.changes, fmt.Sprintf("Duration: %v -> %v", e1.Duration(), e2.Duration()))
	}
	if e1.Track() != e2.Track() {
		e.changes = append(e.changes, fmt.Sprintf("Track: %v -> %v", e1.Track(), e2.Track()))
	}
	if e1.Type() != e2.Type() {
		e.changes = append(e.changes, fmt.Sprintf("Event Type: %v -> %v", e1.Type(), e2.Type()))
	}
	if len(e.changes) == 0 {
		return nil
	}
	return e
}
