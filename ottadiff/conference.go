package ottadiff

import (
	"fmt"
	"github.com/draringi/ottaconf"
)

type conferenceChange struct {
	changes []string
}

func (c *conferenceChange) Type() ChangeType {
	return Modification
}

func (c *conferenceChange) String() string {
	result := "Conference Changes:\n"
	for _, s := range c.changes {
		result += fmt.Sprintf("\t%v\n", s)
	}
	return result
}

func diffConferenceMeta(c1, c2 *ottaconf.Conference) *conferenceChange {
	c := new(conferenceChange)
	if c1.Release() != c2.Release() {
		c.changes = append(c.changes, fmt.Sprintf("Version: %v -> %v", c1.Release(), c2.Release()))
	}
	if c1.Title() != c2.Title() {
		c.changes = append(c.changes, fmt.Sprintf("Title: %v -> %v", c1.Title(), c2.Title()))
	}
	if c1.Subtitle() != c2.Title() {
		c.changes = append(c.changes, fmt.Sprintf("Subtitle: %v -> %v", c1.Subtitle(), c2.Subtitle()))
	}
	if c1.Venue() != c2.Venue() {
		c.changes = append(c.changes, fmt.Sprintf("Venue: %v -> %v", c1.Venue(), c2.Venue()))
	}
	if c1.City() != c2.City() {
		c.changes = append(c.changes, fmt.Sprintf("City: %v -> %v", c1.City(), c2.City()))
	}
	if c1.StartDate() != c2.StartDate() {
		c.changes = append(c.changes, fmt.Sprintf("Start Date: %v -> %v", c1.StartDate(), c2.StartDate()))
	}
	if c1.EndDate() != c2.EndDate() {
		c.changes = append(c.changes, fmt.Sprintf("End Date: %v -> %v", c1.EndDate(), c2.EndDate()))
	}
	return c
}
