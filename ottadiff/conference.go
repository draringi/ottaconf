package ottadiff

import (
	"fmt"
	"github.com/draringi/ottaconf"
	"strconv"
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
		result += "\t" + s + "\n"
	}
	return result
}

func diffConferenceMeta(c1, c2 *ottaconf.Conference) *conferenceChange {
	c := new(conferenceChange)
	if c1.Title() != c2.Title() {
		c.changes = append(c.changes, "Title: "+c1.Title()+" -> "+c2.Title())
	}
	if c1.Subtitle() != c2.Title() {
		c.changes = append(c.changes, "Subtitle: "+c1.Subtitle()+" -> "+c2.Subtitle())
	}
	if c1.Venue() != c2.Venue() {
		c.changes = append(c.changes, "Venue: "+c1.Venue()+" -> "+c2.Venue())
	}
	if c1.City() != c2.City() {
		c.changes = append(c.changes, "City: "+c1.City()+" -> "+c2.City())
	}
	if c1.StartDate() != c2.StartDate() {
		c.changes = append(c.changes, "Start Date: "+c1.StartDate().String()+" -> "+c2.StartDate().String())
	}
	if c1.EndDate() != c2.EndDate() {
		c.changes = append(c.changes, "End Date: "+c1.EndDate().String()+" -> "+c2.EndDate().String())
	}
	return c
}
